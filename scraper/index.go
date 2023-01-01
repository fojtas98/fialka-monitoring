package scraper

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func CreateScraper() (int, *colly.Collector) {
	var c = colly.NewCollector()
	intValue := 0
	c.OnXML("/html/body/header/div[1]/div[2]/div/div[3]/div/div/text()", func(x *colly.XMLElement) {
		v := strings.TrimSpace(x.Text)
		if v != "" {
			intValue, _ = strconv.Atoi(v)
		}
	})
	return intValue, c
}
