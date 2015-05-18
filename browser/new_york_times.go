package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"bytes"
	"errors"
	"net/url"
	"strings"
	"unicode"
	"io"
	goose "github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/advancedlogic/GoOse"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
)

type NewYorkTimes struct {}

func (s *NewYorkTimes) Recognize(link string) bool {
	u, err := url.Parse(link)
	if err != nil {
		return false
	}
	parts := strings.Split(u.Host, ".")
	if len(parts) > 2 {
		parts = parts[1:]
	}
	domain := strings.Join(parts, ".") 
	return domain == "nytimes.com"
}

func (s *NewYorkTimes) Login(bow *browser.Browser, creds map[string]string) error {
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

func (s *NewYorkTimes) Get(bow *browser.Browser, link string) (string, *html.Node, error) {
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

func (e *NewYorkTimes) Extract(permalink string, page *html.Node) (*Post, error) {
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

	return &Post{
		Title: article.Title,
		Permalink: article.CanonicalLink,
		Content: content,
		Images: []string{article.TopImage},
		Tags: strings.Split(article.MetaKeywords, ","),
		Source: "new-york-times",
		Language: article.MetaLang,
	}, nil
}
