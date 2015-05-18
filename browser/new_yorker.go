package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"strings"
	"io"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
)

type NewYorker struct {}

func (s *NewYorker) Recognize(link string) bool {
	return domainMatch(link, "newyorker.com")
}

func (s *NewYorker) Login(bow *browser.Browser, creds map[string]string) (err error) {
	return
}

func (s *NewYorker) Get(bow *browser.Browser, link string) (string, *html.Node, error) {
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

func (s *NewYorker) Extract(permalink string, page *html.Node) (*Post, error) {
	doc := goquery.NewDocumentFromNode(page)
	doc.Find("hgroup h1").Text()

	title := doc.Find("hgroup h1").Text()
	content := doc.Find(".articleBody p").Text()
	topImage, _ := doc.Find(".articleBody figure.featured a").Attr("href")
	author, _ := doc.Find(".author-details meta[itemprop=name]").Attr("content")
	keywords, _ := doc.Find("meta[name=news_keywords]").Attr("content")
	lang, _ := doc.Find("html").Attr("lang")

	return &Post{
		Title: title,
		Permalink: permalink,
		Content: content,
		Images: []string{topImage},
		Authors: []string{author},
		Tags: strings.Split(keywords, ","),
		Source: "new-yorker",
		Language: lang,
	}, nil
}
