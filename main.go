package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/db"
	"github.com/gocolly/colly"
)

func main() {
	ctx := context.Background()

	bunDb := db.ConnectToDB(ctx)
	i := 1
	var c = colly.NewCollector()
	for {
		numberOfPeople := 0
		c.OnXML("/html/body/header/div[1]/div[2]/div/div[3]/div/div/text()", func(x *colly.XMLElement) {
			v := strings.TrimSpace(x.Text)
			if v != "" {
				numberOfPeople, _ = strconv.Atoi(v)
			}
		})
		c.Visit("https://www.nafialce.cz/bazen-sauny")
		fmt.Printf("looping for %d time\n", i)
		ava := &db.Availibility{NumberOfPeople: numberOfPeople, Time: time.Now()}
		bunDb.NewInsert().Model(ava).Exec(ctx)
		time.Sleep(time.Minute * 1)
		i++
	}
}
