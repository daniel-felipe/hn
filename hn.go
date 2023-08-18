package hn

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getDoc() *goquery.Document {
    resp, err := http.Get("https://news.ycombinator.com/newest")
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        panic(err)
    }

    return doc
}

func printNews(el *goquery.Selection) {
    fmt.Println("📰", el.Children().Text())
    link, _ := el.Children().Attr("href")
    fmt.Printf("-> %s\n\n", link)
}

func ShowLatests(total int) {
    doc := getDoc()

    doc.Find(".titleline").EachWithBreak(func (i int, el *goquery.Selection) bool {
        if i == total {
            return false
        }

        printNews(el)

        return true
    });
}

