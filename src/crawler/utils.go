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
	values := make([][]string, 3)
	for i := range values {
		values[i] = make([]string, 0)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www"+DOMAIN, "https://www."+DOMAIN, DOMAIN),
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
				values[r] = append(values[r], StandardizeSpaces(val.Text))
			})
			row.ForEach("td[nowrap]", func(_ int, el *colly.HTMLElement) {
				values[r] = append(values[r], el.ChildAttr("a", "href"))
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
	return createCVEembed(values)
}

func createCVEembed(vals [][]string) *discordgo.MessageEmbed {
	message := embed.NewEmbed()
	message.SetImage(CVE_LOGO_URL)
	message.SetTitle("CVEs")
	message.SetDescription("Recent CVEs with score >=3")
	message.SetThumbnail(CVE_THUMBNAIL_URL)
	message.SetColor(0xffff00)
	for _, cve := range vals {
		fieldTitle := cve[CVE_ID]
		fieldDesc := ":lady_beetle: Vuln type: " + cve[VULN_TYPE] + "\n:scales: Score: " + cve[SCORE] + "\n:link: [link](" + DOMAIN + cve[len(cve)-1] + ")"
		message.AddField(fieldTitle, fieldDesc)
	}
	return message.MessageEmbed
}
