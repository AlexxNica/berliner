package scrape

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/content"
	"strings"
)

type newYorker struct{}

func (s *newYorker) recognize(url string) bool {
	return isDomain(url, "newyorker.com")
}

func (s *newYorker) scrape(page *html.Node) (content.Post, error) {
	doc := goquery.NewDocumentFromNode(page)
	return content.Post{
		Title:    doc.Find("hgroup h1").Text(),
		Body:     must(doc.Find(".articleBody p").WrapAllHtml("<div></div>").Html()),
		Authors:  []string{doc.Find(".author-details meta[itemprop=name]").AttrOr("content", "")},
		Tags:     strings.Split(doc.Find("meta[name=news_keywords]").AttrOr("content", ""), ","),
		Source: "New Yorker",
		Language: doc.Find("html").AttrOr("lang", "en"),
	}, nil
}

func init() {
	register(&newYorker{})
}
