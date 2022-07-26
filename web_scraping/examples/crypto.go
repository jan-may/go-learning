package examples

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func Crypto() {
	fName := "cryptocoinmarketcap.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML(".cmc-table__table-wrapper-outer div table tbody tr", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText(".cmc-table__cell-name"))
		writer.Write([]string{
			e.ChildText(".cmc-table__cell--sort-by__name"),
			e.ChildText(".col-symbol"),
			e.ChildAttr("a.price", "data-usd"),
			e.ChildAttr("a.volume", "data-usd"),
			// e.ChildAttr(".market-cap", "data-usd"),
			// e.ChildText(".percent-1h"),
			// e.ChildText(".percent-24h"),
			// e.ChildText(".percent-7d"),
		})
		if err := writer.Error(); err != nil {
			log.Fatalln("error writing to csv:", err)
		}
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for results\n", fName)
}