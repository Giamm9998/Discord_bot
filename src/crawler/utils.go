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
	cats := make([]string, 0)

	// slice 2d for the CVES
	vals := make([][]string, 5)
	for i := range vals {
		vals[i] = make([]string, 0)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.cvedetails.com", "cvedetails.com"),
	)

	//TODO: improve css selectors
	c.OnHTML("table[id=vulnslisttable]", func(e *colly.HTMLElement) {
		e.ForEach("tbody tr th", func(_ int, category *colly.HTMLElement) {
			// gets all the category names
			cats = append(cats, StandardizeSpaces(category.Text))
		})
		e.ForEachWithBreak(".srrowns", func(r int, row *colly.HTMLElement) bool {
			if r == 5 {
				return false
			}
			row.ForEach("td", func(_ int, val *colly.HTMLElement) {
				vals[r] = append(vals[r], StandardizeSpaces(val.Text))
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
	return createTable(cats, vals)
}

func createTable(cats []string, vals [][]string) *discordgo.MessageEmbed {
	ascii_table := "\n+=====+===============+=========+\n| "
	separator := "\n-------------------------------------------------\n"
	cells := [3]string{cats[CVE_ID], cats[SCORE], cats[VULN_TYPE]}

	for _, cell := range cells {
		ascii_table += cell + " | "
	}
	ascii_table += "\n+=====+===============+=========+\n"

	for _, row := range vals {
		cells = [3]string{row[CVE_ID], row[SCORE], row[VULN_TYPE]}
		for _, col := range cells {
			ascii_table += col + " | "
		}
		ascii_table += separator
	}
	ascii_table += "\n+=====+===================+=========+"

	return create_embed(ascii_table)
}

func create_embed(descrption string) *discordgo.MessageEmbed {
	embd := embed.NewEmbed()
	embd.SetImage(CVE_LOGO_URL)
	embd.SetTitle("CVEs")
	embd.SetDescription(descrption)
	embd.SetThumbnail(CVE_THUMBNAIL_URL)
	return embd.MessageEmbed
}
