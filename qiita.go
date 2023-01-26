package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("article > h2", func(h *colly.HTMLElement) {
		title := h.DOM.Find("a").Text()
		url   := h.ChildAttr("a", "href")
		fmt.Println("title", title)
		fmt.Println("url", url)
	})

	c.Visit("https://qiita.com/")
}
