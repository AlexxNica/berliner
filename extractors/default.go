package extractors

import (
	"bytes"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	goose "github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/advancedlogic/GoOse"
)

type Default struct {
	post *Post
}

func (e *Default) Recognize(link string) bool {
	return true
}

func (e *Default) SetPost(post *Post) {
	e.post = post
}

func (e *Default) Get() (*html.Node, error) {
	resp, err := http.Get(e.post.Permalink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	e.post.Permalink = resp.Request.URL.String()
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
	article := g.ExtractFromRawHtml(e.post.Permalink, rawHtml)
	content := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}
		return r
	}, article.CleanedText)

	source := ""
	url, err := url.Parse(e.post.Permalink)
	if err == nil {
		source = url.Host
	}

	e.post.Title = article.Title
	e.post.Permalink = article.CanonicalLink
	e.post.Content = content
	e.post.Images = []string{article.TopImage}
	e.post.Tags = strings.Split(article.MetaKeywords, ",")
	e.post.Source = source
	e.post.Language = article.MetaLang
	return e.post, nil
}
