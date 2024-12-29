package main

import (
	"fmt"
	"strings"

	colly "github.com/gocolly/colly"
)

func main() {
	
	c := colly.NewCollector()

	c.OnHTML("div.evenPerson", func(e *colly.HTMLElement) {
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

	c.Visit("https://yazd.ac.ir/faculties/engineering/departments/computer/people/academics")
}