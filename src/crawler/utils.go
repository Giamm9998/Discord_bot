package crawler

import (
	"fmt"
	"log"
	"strings"

	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
)

func StandardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func getCVE() *discordgo.MessageEmbed {

	// the assignments inside OnHTML works only with slices. Why???
	// cats := make([]string, 0)

	// slice 2d for the CVES
	vals := make([][]string, 3)
	for i := range vals {
		vals[i] = make([]string, 0)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.cvedetails.com", "cvedetails.com", DOMAIN),
	)

	//TODO: improve css selectors
	c.OnHTML("table[id=vulnslisttable]", func(e *colly.HTMLElement) {
		// e.ForEach("tbody tr th", func(_ int, category *colly.HTMLElement) {
		// 	// gets all the category names
		// 	cats = append(cats, StandardizeSpaces(category.Text))
		// })
		//TODO: take only recent Weekly/Daily CVE
		e.ForEachWithBreak(".srrowns", func(r int, row *colly.HTMLElement) bool {
			if r == 3 {
				return false
			}
			row.ForEach("td", func(_ int, val *colly.HTMLElement) {
				vals[r] = append(vals[r], StandardizeSpaces(val.Text))
			})
			row.ForEach("td[nowrap]", func(_ int, el *colly.HTMLElement) {
				vals[r] = append(vals[r], el.ChildAttr("a", "href"))
			})
			return true
		})
	})

	err := c.Visit(URL)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("DONE!")
	return createCVEembed(vals)
}

func createCVEembed(vals [][]string) *discordgo.MessageEmbed {
	embd := embed.NewEmbed()
	embd.SetImage(CVE_LOGO_URL)
	embd.SetTitle("CVEs")
	embd.SetDescription("Recent CVEs with score >=3")
	embd.SetThumbnail(CVE_THUMBNAIL_URL)
	embd.SetColor(0xffff00)
	for _, cve := range vals {
		field_title := cve[CVE_ID]
		field_desc := ":lady_beetle: Vuln type: " + cve[VULN_TYPE] + "\n:scales: Score: " + cve[SCORE] + "\n:link: [link](" + DOMAIN + cve[len(cve)-1] + ")"
		embd.AddField(field_title, field_desc)
	}
	return embd.MessageEmbed
}
