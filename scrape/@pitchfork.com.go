package scrape

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/content"
)

type pitchfork struct{}

func (s *pitchfork) recognize(url string) bool {
	return isDomain(url, "pitchfork.com")
}

func (s *pitchfork) scrape(page *html.Node) (content.Post, error) {
	doc := goquery.NewDocumentFromNode(page)
	infoDiv := doc.Find("#content #Main .info")
	body, _ := doc.Find(".object-detail .editorial").Html()

	return content.Post{
		Title:    infoDiv.Find("h1").Text() + ": " + infoDiv.Find("h2").Text(), // artist + album name
		Body:     body,
		Authors:  []string{infoDiv.Find("h4 a").Text()}, // TODO this should work but doesn't, debug
		Tags:     nil,
		Origin:   "Pitchfork",
		Language: doc.Find("html").AttrOr("lang", "en"),
		Images:   []string{doc.Find("#content #main .artwork img").AttrOr("src", "")},
	}, nil
}

func init() {
	register(&pitchfork{})
}
