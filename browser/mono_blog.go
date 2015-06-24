package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"io"
)

type monoBlog struct{}

func (s *monoBlog) slug() string {
	return "mono-blog"
}

func (s *monoBlog) recognize(link string) bool {
	return domainMatch(link, "mono-blog.com")
}

func (s *monoBlog) login(bow *browser.Browser, creds map[string]string) error {
	return nil
}

func (s *monoBlog) get(bow *browser.Browser, link string) (string, *html.Node, error) {
	err := bow.Open(link)
	if err != nil {
		return "", nil, err
	}
	r, w := io.Pipe()
	go func() {
		_, _ = bow.Download(w)
		w.Close()
	}()
	r2, err := charset.NewReader(r, bow.ResponseHeaders().Get("content-type"))
	if err != nil {
		return "", nil, err
	}
	page, err := html.Parse(r2)
	if err != nil {
		return "", nil, err
	}
	return bow.Url().String(), page, nil
}

func (s *monoBlog) extract(permalink string, page *html.Node) (*Post, error) {
	doc := goquery.NewDocumentFromNode(page)

	title := doc.Find("h2.entry-title").Text()
	content, _ := doc.Find(".entry-content").Html()
	author := doc.Find(".vcard .fn.n").Text()
	tags := []string{}
	doc.Find(".entry-category a[rel~=\"category\"]").Each(func(i int, s *goquery.Selection) {
		tags = append(tags, s.Text())
	})
	lang, _ := doc.Find("html").Attr("lang")

	p := &Post{
		Title:     title,
		Content:   content,
		Permalink: permalink,
		Authors:   []string{author},
		Tags:      tags,
		Source:    s.slug(),
		Language:  lang,
	}
	return p, nil
}

func init() {
	register(&monoBlog{})
}
