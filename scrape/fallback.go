package scrape

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"

	"github.com/thatguystone/swan"
	"github.com/s3ththompson/berliner/content"
)

type fallback struct{}

func (s *fallback) recognize(url string) bool {
	return true
}

func (s *fallback) scrape(page *html.Node) (content.Post, error) {
	var rawHTML bytes.Buffer
	err := html.Render(&rawHTML, page)
	if err != nil {
		return content.Post{}, err
	}
	article, err := swan.FromHTML("", rawHTML.Bytes()) // TODO: why does it need url??
	if err != nil {
		return content.Post{}, err
	}
	body, err := article.TopNode.Html()
	if err != nil {
		return content.Post{}, err
	}
	var tags []string
	if article.Meta.Keywords != "" {
		tags = strings.Split(article.Meta.Keywords, ",")
	} else {
		tags = make([]string, 0)
	}
	p := content.Post{
		Title:     article.Meta.Title,
		Body:      body,
		Permalink: article.Meta.Canonical,
		Tags:      tags,
		Language:  article.Meta.Lang,
	}
	return p, nil
}
