package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func Parse(cmd *cobra.Command, args []string) {
	c := make(chan *html.Node)
	p1 := &Pipe{
		workers: 20,
		do:      get,
		in:      readLines(),
		out:     c,
	}
	err := p1.pipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	articles := make(chan *Article)
	p2 := &Pipe{
		workers: 5,
		do:      parse,
		in:      c,
		out:     articles,
	}
	err = p2.pipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	for article := range articles {
		fmt.Fprintln(os.Stdout, article.title)
	}
}

func get(link string, out chan *html.Node) {
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

func parse(page *html.Node, out chan *Article) {
	doc := goquery.NewDocumentFromNode(page)
	out <- &Article{
		title: doc.Find("head title").Text(),
	}
}
