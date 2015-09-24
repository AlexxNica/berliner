package scrape

import (
	"io"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html/charset"
	"github.com/s3ththompson/berliner/content"
)

type Client struct {
	Bow *browser.Browser
}

func NewClient() *Client {
	c := &Client{}
	c.init()
	return c
}

func (c *Client) init() {
	c.Bow = surf.NewBrowser()
	c.Bow.AddRequestHeader("Accept", "text/html")
	c.Bow.AddRequestHeader("Accept-Charset", "utf8")
}

func (c *Client) Get(url string) (*html.Node, string, error) {
	err := c.Bow.Open(url)
	if err != nil {
		return nil, "", err
	}
	r, w := io.Pipe()
	go func() {
		_, _ = c.Bow.Download(w)
		w.Close()
	}()
	r2, err := charset.NewReader(r, c.Bow.ResponseHeaders().Get("content-type"))
	if err != nil {
		return nil, "", err
	}
	page, err := html.Parse(r2)
	if err != nil {
		return nil, "", err
	}
	return page, c.Bow.Url().String(), nil
}

func (c *Client) GetPost(url string) (content.Post, error) {
	page, permalink, err := c.Get(url)
	if err != nil {
		return content.Post{}, err
	}
	s := scrapers.byURL(permalink)
	post, err := s.scrape(page)
	if err != nil {
		return content.Post{}, err
	}
	post.Permalink = permalink
	if post.Origin == "" {
		post.Origin = getDomain(permalink)
	}
	return post, nil
}
