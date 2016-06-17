package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/content"
	"golang.org/x/net/html"
)

type pitchfork struct{}

func (s *pitchfork) recognize(url string) bool {
	return isDomain(url, "pitchfork.com")
}

func (s *pitchfork) scrape(page *html.Node) (content.Post, error) {
	doc := goquery.NewDocumentFromNode(page)
	infoDiv := doc.Find("#content #main .info")
	article, _ := doc.Find(".object-detail .editorial").Html()
	score := infoDiv.Find("span.score").Text()
	body := "<h2>Score: " + score + "</h2>" + article

	return content.Post{
		Title:    infoDiv.Find("h1").First().Text() + ": " + infoDiv.Find("h2").Text(), // artist + album name
		Body:     body,
		Authors:  []string{infoDiv.Find("h4 a").Text()},
		Tags:     nil,
		Origin:   "Pitchfork",
		Language: doc.Find("html").AttrOr("lang", "en"),
		Images:   []content.Image{content.NewImage(doc.Find("#content #main .artwork img").AttrOr("src", ""), "")},
	}, nil
}

func init() {
	register(&pitchfork{})
}
