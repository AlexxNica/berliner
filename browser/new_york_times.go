package browser

import (
	"errors"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"io"
	"strings"
)

type newYorkTimes struct{}

func (s *newYorkTimes) slug() string {
	return "new-york-times"
}

func (s *newYorkTimes) recognize(link string) bool {
	return domainMatch(link, "nytimes.com")
}

func (s *newYorkTimes) login(bow *browser.Browser, creds map[string]string) error {
	err := bow.Open("https://myaccount.nytimes.com/auth/login")
	if err != nil {
		return err
	}
	fm, err := bow.Form("form.loginForm")
	if err != nil {
		return err
	}
	userid, ok1 := creds["userid"]
	password, ok2 := creds["password"]
	if !ok1 || !ok2 {
		return errors.New("New York Times credentials require `userid` and `password`.")
	}
	fm.Input("userid", userid)
	fm.Input("password", password)
	if err := fm.Submit(); err != nil {
		return err
	}
	return nil
}

func (s *newYorkTimes) get(bow *browser.Browser, link string) (string, *html.Node, error) {
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

func (s *newYorkTimes) extract(permalink string, page *html.Node) (*Post, error) {
	doc := goquery.NewDocumentFromNode(page)

	title, _ := doc.Find("meta[name=hdl]").Attr("content")
	content := doc.Find("p.story-body-text.story-content").Text()
	topImage, _ := doc.Find(".lede-container figure .image img").Attr("data-mediaviewer-src")
	author, _ := doc.Find("meta[name=author]").Attr("content")
	keywords, _ := doc.Find("meta[name=keywords]").Attr("content")
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
	register(&newYorkTimes{})
}
