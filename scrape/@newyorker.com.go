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
		Body:     strings.Join(doc.Find(".articleBody > p").Map(func(_ int, s *goquery.Selection) string {
			html, _ := s.Html()
			return "<p>" + html + "</p>"
		}), "\n"),
		Authors:  []string{doc.Find(".author-details meta[itemprop=name]").AttrOr("content", "")},
		Tags:     strings.Split(doc.Find("meta[name=news_keywords]").AttrOr("content", ""), ","),
		Source:   "New Yorker",
		Language: doc.Find("html").AttrOr("lang", "en"),
		Images:   []string{doc.Find("figure.featured img").AttrOr("src", "")},
	}, nil
}

func init() {
	register(&newYorker{})
}
