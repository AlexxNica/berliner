package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
)

type Strategy interface {
	recognize(string) bool
	login(*browser.Browser, map[string]string) error
	get(*browser.Browser, string) (string, *html.Node, error)
	extract(string, *html.Node) (*Post, error)
}

type StrategyList struct {
	strats map[string]Strategy
	fallback   Strategy
}

func (sl *StrategyList) findMatch(link string) Strategy {
	for _, s := range sl.strats {
		if s.recognize(link) {
			return s
		}
	}
	return sl.fallback
}