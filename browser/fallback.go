package browser

import (
	"bytes"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"net/http"
	"net/url"
	"strings"

	goose "github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/advancedlogic/GoOse"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
)

type fallback struct{}

func (s *fallback) slug() string {
	return ""
}

func (s *fallback) recognize(link string) bool {
	return true
}

func (s *fallback) login(bow *browser.Browser, creds map[string]string) error {
	return nil
}

func (s *fallback) get(bow *browser.Browser, link string) (string, *html.Node, error) {
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

func (s *fallback) extract(permalink string, page *html.Node) (*Post, error) {
	var raw bytes.Buffer
	err := html.Render(&raw, page)
	if err != nil {
		return nil, err
	}
	rawHtml := raw.String()
	g := goose.New()
	article := g.ExtractFromRawHtml(permalink, rawHtml)

	source := ""
	u, err := url.Parse(permalink)
	if err == nil {
		source = u.Host
	}

	p := &Post{
		Title:     article.Title,
		Content:   article.CleanedText,
		Permalink: article.CanonicalLink,
		Tags:      strings.Split(article.MetaKeywords, ","),
		Source:    source,
		Language:  article.MetaLang,
	}
	p.addImage(article.TopImage, "")
	return p, nil
}
