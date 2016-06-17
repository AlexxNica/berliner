package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/content"
	"golang.org/x/net/html"
)

type itsnicethat struct{}

func (s *itsnicethat) recognize(url string) bool {
	return isDomain(url, "itsnicethat.com")
}

func (s *itsnicethat) scrape(page *html.Node) (content.Post, error) {
	doc := goquery.NewDocumentFromNode(page)
	body, _ := doc.Find("article .text-slice.container .text-slice__text").Html()
	post := content.Post{
		Title:    doc.Find("article header.article-header-slice container h1").Text(),
		Body:     body,
		Authors:  []string{doc.Find("article .article-header-slice__byline a").First().Text()},
		Tags:     []string{},
		Origin:   "It's Nice That",
		Language: doc.Find("html").AttrOr("lang", "en"),
	}
	doc.Find(".article-header-slice__breadcrumbs a").Each(func(_ int, s *goquery.Selection) {
		tag := s.Text()
		post.Tags = append(post.Tags, tag)
	})
	doc.Find(".images-slice.container figure").Each(func(_ int, s *goquery.Selection) {
		src := s.Find("img.images-slice__image").AttrOr("src", "")
		caption, _ := doc.Find("figcaption.images-slice__caption").Html()
		post.AddImage(src, caption)
	})
	return post, nil
}

func init() {
	register(&itsnicethat{})
}
