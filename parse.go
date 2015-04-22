package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"net/http"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func Parse(cmd *cobra.Command, args []string) {
	links := readLines()
	pages := get(links)
	articles := parse(pages)
	for article := range articles {
		fmt.Fprintln(os.Stdout, article.title)
	}
}

func get(links <-chan string) <-chan *html.Node {
	out := make(chan *html.Node)
	go func() {
		var wg sync.WaitGroup
		for link := range links {
			wg.Add(1)
			go func(link string) {
				defer wg.Done()
				innerGet(link, out)
			}(link)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func innerGet(link string, out chan *html.Node) {
	resp, err := http.Get(link)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Read error: %s\n", link)
		return
	}
	r, err := charset.NewReader(resp.Body, resp.Header.Get("content-type"))
	if err != nil {
		fmt.Printf("UTF8 error: %s\n", link)
		return
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Printf("Parse error: %s\n", link)
		return
	}
	out <- doc
}

type Article struct {
	title string
}

func parse(pages <-chan *html.Node) <-chan *Article {
	out := make(chan *Article)
	go func() {
		var wg sync.WaitGroup
		for page := range pages {
			wg.Add(1)
			go func(page *html.Node) {
				defer wg.Done()
				innerParse(page, out)
			}(page)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func innerParse(page *html.Node, out chan *Article) {
	doc := goquery.NewDocumentFromNode(page)
	out <- &Article{
		title: doc.Find("head title").Text(),
	}
}
