package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"net/http"
)

type tlm struct{}

func (s *tlm) slug() string {
	return "tenderlovemaking"
}

func (s *tlm) recognize(link string) bool {
	return domainMatch(link, "tenderlovemaking.com")
}

func (s *tlm) login(bow *browser.Browser, creds map[string]string) error {
	return nil
}

func (s *tlm) get(bow *browser.Browser, link string) (string, *html.Node, error) {
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

func (s *tlm) extract(permalink string, page *html.Node) (*Post, error) {
	doc := goquery.NewDocumentFromNode(page)

	title := doc.Find("h1.posted-title").Text()
	content, _ := doc.Find(".entry").Html()
	author := "Aaron Patterson" // It's a single-author blog
	tags := []string{}

	p := &Post{
		Title:     title,
		Content:   content,
		Permalink: permalink,
		Authors:   []string{author},
		Tags:      tags,
		Source:    s.slug(),
	}
	return p, nil
}

func init() {
	register(&tlm{})
}
