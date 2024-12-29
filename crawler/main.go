package main

import (
	"fmt"
	"strings"

	colly "github.com/gocolly/colly"
)

func main() {
	
	c := colly.NewCollector()
	d := colly.NewCollector()
	f := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "departments") {
			d.Visit(link)
		}
	})

	d.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "academics") {
			fmt.Println(link)
			f.Visit(link)
		}
	})

	f.OnHTML("div.evenPerson", func(e *colly.HTMLElement) {
		name := e.ChildAttr("a","title")
		if name != "" {
			fmt.Printf("professor name: %s\n", name)
		}
		if strings.Contains(e.Text, "[at]") {
			fmt.Printf("email found: %q\n", e.Text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Received:", string(r.Body))
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	d.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	d.OnResponse(func(r *colly.Response) {
		// fmt.Println("Received:", string(r.Body))
	})

	d.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.Visit("https://yazd.ac.ir/")
}