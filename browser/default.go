package browser

import (
	"bytes"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	goose "github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/advancedlogic/GoOse"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
)

type Default struct {}

func (s *Default) recognize(link string) bool {
	return true
}

func (s *Default) login(bow *browser.Browser, creds map[string]string) error {
	return nil
}

func (s *Default) get(bow *browser.Browser, link string) (string, *html.Node, error) {
	resp, err := http.Get(link)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()
	permalink := resp.Request.URL.String()
	r, err := charset.NewReader(resp.Body, resp.Header.Get("content-type"))
	if err != nil {
		return "", nil, err
	}
	page, err := html.Parse(r)
	if err != nil {
		return "", nil, err
	}
	return permalink, page, nil
}

func (s *Default) extract(permalink string, page *html.Node) (*Post, error) {
	var raw bytes.Buffer
	err := html.Render(&raw, page)
	if err != nil {
		return nil, err
	}
	rawHtml := raw.String()
	g := goose.New()
	article := g.ExtractFromRawHtml(permalink, rawHtml)
	content := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}
		return r
	}, article.CleanedText)

	source := ""
	u, err := url.Parse(permalink)
	if err == nil {
		source = u.Host
	}

	return &Post{
		Title: article.Title,
		Permalink: article.CanonicalLink,
		Content: content,
		Images: []string{article.TopImage},
		Tags: strings.Split(article.MetaKeywords, ","),
		Source: source,
		Language: article.MetaLang,
	}, nil
}
