package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/content"
	"golang.org/x/net/html"
)

type newYorker struct{}

func (s *newYorker) recognize(url string) bool {
	return isDomain(url, "newyorker.com")
}

func (s *newYorker) scrape(page *html.Node) (content.Post, error) {
	doc := goquery.NewDocumentFromNode(page)
	src := doc.Find("figure.featured a").AttrOr("href", "")
	caption, _ := doc.Find("figure.featured figcaption .caption-text").Html()
	return content.Post{
		Title: doc.Find("hgroup h1").Text(),
		Body: strings.Join(doc.Find(".articleBody > p").Map(func(_ int, s *goquery.Selection) string {
			html, _ := s.Html()
			return "<p>" + html + "</p>"
		}), "\n"),
		Authors:  []string{doc.Find(".author-details meta[itemprop=name]").AttrOr("content", "")},
		Tags:     strings.Split(doc.Find("meta[name=news_keywords]").AttrOr("content", ""), ","),
		Origin:   "The New Yorker",
		Language: doc.Find("html").AttrOr("lang", "en"),
		Images:   []content.Image{content.NewImage(src, caption)},
	}, nil
}

func init() {
	register(&newYorker{})
}
