package main

import (
	"context"
	"fmt"
	"time"

	"example.com/db"
	"example.com/scraper"
)

func main() {
	ctx := context.Background()

	bunDb := db.ConnectToDB(ctx)
	numberOfPeople, c := scraper.CreateScraper()
	i := 1
	for {
		c.Visit("https://www.nafialce.cz/bazen-sauny")
		fmt.Printf("looping for %d time\n", i)
		ava := &db.Availibility{NumberOfPeople: numberOfPeople, Time: time.Now()}
		bunDb.NewInsert().Model(ava).Exec(ctx)
		time.Sleep(time.Minute * 1)
		i++
	}
}
