package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/gocolly/colly"
)

type Laptop struct {
	Name string
	Price string
}

func JumiaLaptops() {
	fName := "Laptops.csv"
	file, err := os.Create(fName);

	if err != nil {
		log.Fatal("Couldn't create file, err: %q", err)
		return
	}

	defer file.Close();

	writer := csv.NewWriter(file);
	defer writer.Flush()

	c := colly.NewCollector()

	c.OnRequest(func (r *colly.Request){
		fmt.Println("Visiting ", r.URL);
	})

	c.OnHTML(".prd._fb", func(e *colly.HTMLElement) {
		laptop := Laptop{};
		laptop.Name = e.DOM.Find("a").Find(".info").Find(".name").Text();
		laptop.Price = e.DOM.Find("a").Find(".info").Find(".prc").Text();
		row := []string{laptop.Name, laptop.Price};
		writer.Write(row);
	})


	c.Visit("https://www.jumia.co.ke/laptops/");
}



func main() {
	scrape_scheduler := gocron.NewScheduler(time.UTC);
	scrape_scheduler.Every(2).Minute().Do(JumiaLaptops);
	scrape_scheduler.StartBlocking();
}

