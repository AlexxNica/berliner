package extractors

import (
	"bytes"
	"net/http"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"strings"
	"unicode"

	goose "github.com/advancedlogic/GoOse"
)

type Default struct {
	link string
}

func (e *Default) Recognize(link string) bool {
	return true
}

func (e *Default) SetLink(link string) {
	e.link = link
}

func (e *Default) Get() (*html.Node, error) {
	resp, err := http.Get(e.link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	e.link = resp.Request.URL.String()
	r, err := charset.NewReader(resp.Body, resp.Header.Get("content-type"))
	if err != nil {
		return nil, err
	}
	page, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (e *Default) Extract(page *html.Node) (*Post, error) {
	var raw bytes.Buffer
	err := html.Render(&raw, page)
	if err != nil {
		return nil, err
	}
	rawHtml := raw.String()
	g := goose.New()
	article := g.ExtractFromRawHtml(e.link, rawHtml)
	content := strings.Map(func(r rune) rune {
        if unicode.IsSpace(r) {
            return ' '
        }
        return r
    }, article.CleanedText)

	return &Post{
		Title: article.Title,
		Content: content,
		Link: article.CanonicalLink,
		Image: article.TopImage,
	}, nil
}