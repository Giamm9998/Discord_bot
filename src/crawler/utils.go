package crawler

import (
	"fmt"
	"log"

	"newsBot/db"
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

	// slice 2d for the CVES
	values := make([][]string, 3)
	for i := range values {
		values[i] = make([]string, 0)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("https://"+DOMAIN, DOMAIN, "www."+DOMAIN),
	)

	c.OnHTML("ul[id=latestVulns]", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("li", func(r int, row *colly.HTMLElement) bool {
			if r == 3 {
				return false
			}
			row.ForEach(".col-lg-9 p strong", func(_ int, el *colly.HTMLElement) {
				values[r] = append(values[r], StandardizeSpaces(el.Text))
				values[r] = append(values[r], el.ChildAttr("a", "href"))

			})
			row.ForEach(".col-lg-9  p", func(_ int, el *colly.HTMLElement) {
				values[r] = append(values[r], StandardizeSpaces(el.Text))
			})
			row.ForEach(".col-lg-3", func(_ int, el *colly.HTMLElement) {
				values[r] = append(values[r], StandardizeSpaces(el.Text))
			})
			return true
		})
	})

	err := c.Visit("https://" + DOMAIN)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("DONE!")

	cveData := [][4]string{}
	for _, v := range values {
		cveData = append(cveData, [4]string{v[CVE_ID], strings.Replace(v[VULN_TYPE], v[CVE_ID]+" - ", "", 1), v[SCOREv2], "https://" + DOMAIN + v[LINK]})
	}

	db.Connect()
	db.Write_CVE(cveData)
	db.Disconnect()

	return createCVEembed(values)
}

func createCVEembed(vals [][]string) *discordgo.MessageEmbed {
	message := embed.NewEmbed()
	message.SetImage(CVE_LOGO_URL)
	message.SetTitle("CVEs")
	message.SetDescription("Recent scored CVEs from Nist NVD")
	message.SetThumbnail(CVE_THUMBNAIL_URL)
	message.SetColor(0xffff00)
	for _, cve := range vals {
		fieldTitle := cve[CVE_ID]
		//remove cve id from desc
		cve[VULN_TYPE] = strings.Replace(cve[VULN_TYPE], cve[CVE_ID]+" - ", "", 1)
		fieldDesc := ":lady_beetle: **Vuln description**: " + cve[VULN_TYPE] + "\n:scales: **Score**: " + cve[SCOREv2] + "\n:link: [link](https://" + DOMAIN + cve[LINK] + ")"
		message.AddField(fieldTitle, fieldDesc)
	}
	return message.MessageEmbed
}
