package parcer

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"os"
)

type Laptop struct {
	URL      string
	OldPrice string
	Price    string
	Name     string
}

func Parce(domain string) error {

	lap := Laptop{}
	laptops := []Laptop{}

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.199 Safari/537.36"),
	)

	c.OnHTML("article._uCfbZp._6Lpe9P", func(e *colly.HTMLElement) {
		// Парсимо назву товару
		lap.Name = e.ChildText("h2._-fSawu a")
		// Парсимо ціну товару
		lap.Price = e.ChildText("div._tqVytn._Bfr6tl data")
		// Парсимо стару ціну товару (якщо вона є)
		lap.OldPrice = e.ChildText("div._OHybQh s._0pVbDF data")

	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("scrap.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		headers := []string{"URL", "OldPrice", "Price", "Name"}
		writer.Write(headers)
		for _, lap = range laptops {
			record := []string{
				lap.URL,
				lap.OldPrice,
				lap.Price,
				lap.Name,
			}
			writer.Write(record)
		}
		defer writer.Flush()
	})

	err := c.Visit("https://epicentrk.ua/ua/shop/morozilnye-kamery-i-lari/")
	fmt.Printf("lala %v\n", err)
	return err

}
