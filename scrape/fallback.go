package scrape

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"

	goose "github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/advancedlogic/GoOse" // TODO: update goose
	"github.com/s3ththompson/berliner/content"
)

type fallback struct{}

func (s *fallback) recognize(url string) bool {
	return true
}

// Converts an article with newline-delimited paragraphs to an HTML string
// with each paragraph wrapped in a <p> tag
func textToHtml(text string) string {
	var buf bytes.Buffer
	var t = template.Must(template.New("name").Parse("{{range .}}<p>{{.}}</p>{{end}}"))
	err := t.Execute(&buf, strings.Split(text, "\n\n"))
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (s *fallback) scrape(page *html.Node) (content.Post, error) {
	var raw bytes.Buffer
	err := html.Render(&raw, page)
	if err != nil {
		return content.Post{}, err
	}
	rawHtml := raw.String()
	g := goose.New()
	article := g.ExtractFromRawHtml("", rawHtml) // TODO: why does it need url??
	var tags []string
	if article.MetaKeywords != "" {
		tags = strings.Split(article.MetaKeywords, ",")
	} else {
		tags = make([]string, 0)
	}

	p := content.Post{
		Title:     article.Title,
		Body:      textToHtml(article.CleanedText),
		Permalink: article.CanonicalLink,
		Tags:      tags,
		Language:  article.MetaLang,
	}
	return p, nil
}
