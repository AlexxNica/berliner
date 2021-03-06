package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/content"
	"golang.org/x/net/html"
)

type grantland struct{}

func (s *grantland) recognize(url string) bool {
	return isDomain(url, "grantland.com")
}

func (s *grantland) scrape(page *html.Node) (content.Post, error) {
	doc := goquery.NewDocumentFromNode(page)
	body, _ := doc.Find(".article-body").Html()

	post := content.Post{
		Title:    doc.Find("header.title-card h1.title").Text(),
		Summary:  doc.Find("header.title-card p.summary").Text(),
		Body:     body,
		Authors:  []string{doc.Find(".byline a[rel='author']").First().Text()},
		Tags:     []string{},
		Origin:   "Grantland",
		Language: doc.Find("html").AttrOr("lang", "en"),
	}
	doc.Find("span.feature img").Each(func(_ int, s *goquery.Selection) {
		src := s.AttrOr("src", "")
		post.AddImage(src, "")
	})
	return post, nil
}

func init() {
	register(&grantland{})
}
