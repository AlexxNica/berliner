package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
)

type strategy interface {
	slug() string
	recognize(string) bool
	login(*browser.Browser, map[string]string) error
	get(*browser.Browser, string) (string, *html.Node, error)
	extract(string, *html.Node) (*Post, error)
}

type lookup struct {
	table    map[string]strategy
	fallback strategy
}

func (l *lookup) bySlug(slug string) (s strategy, ok bool) {
	s, ok = l.table[slug]
	return
}

func (l *lookup) byLink(link string) strategy {
	for _, s := range l.table {
		if s.recognize(link) {
			return s
		}
	}
	return l.fallback
}

var strategies lookup = lookup{
	table:    make(map[string]strategy),
	fallback: &fallback{},
}

func register(s strategy) {
	slug := s.slug()
	if _, dup := strategies.table[slug]; dup {
		return
	}
	strategies.table[slug] = s
}
