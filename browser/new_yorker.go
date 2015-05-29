package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"io"
	"strings"
)

type newYorker struct{}

func (s *newYorker) slug() string {
	return "new-yorker"
}

func (s *newYorker) recognize(link string) bool {
	return domainMatch(link, "newyorker.com")
}

func (s *newYorker) login(bow *browser.Browser, creds map[string]string) (err error) {
	return
}

func (s *newYorker) get(bow *browser.Browser, link string) (string, *html.Node, error) {
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

func (s *newYorker) extract(permalink string, page *html.Node) (*Post, error) {
	doc := goquery.NewDocumentFromNode(page)

	title := doc.Find("hgroup h1").Text()
	content := doc.Find(".articleBody p").Text()
	topImage, _ := doc.Find(".articleBody figure.featured a").Attr("href")
	author, _ := doc.Find(".author-details meta[itemprop=name]").Attr("content")
	keywords, _ := doc.Find("meta[name=news_keywords]").Attr("content")
	lang, _ := doc.Find("html").Attr("lang")

	p := &Post{
		title:     title,
		permalink: permalink,
		authors:   []string{author},
		tags:      strings.Split(keywords, ","),
		source:    s.slug(),
		language:  lang,
	}
	p.setContent(content)
	p.addImage(topImage, "")
	return p, nil
}

func init() {
	register(&newYorker{})
}
