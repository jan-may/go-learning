package examples

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Book struct {
	Title string
	Price string
}

func Books() {
	file, err := os.Create("export.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Title", "Price"}
	writer.Write(headers)

	c := colly.NewCollector(colly.AllowedDomains("books.toscrape.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println(r.StatusCode)
	// })

	// c.OnHTML("title", func(e *colly.HTMLElement) {
	// 	fmt.Println(e.Text)
	// })

	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		book := Book{}
		book.Title = e.ChildAttr(".image_container img", "alt")
		book.Price = e.ChildText(".price_color")
		row := []string{book.Title, book.Price}
		writer.Write(row)
	})

	c.OnHTML(".next > a", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	c.Visit("https://books.toscrape.com/")

	fmt.Println("Done")
}