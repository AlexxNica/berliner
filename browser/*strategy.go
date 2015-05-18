package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
)

type Strategy interface {
	Recognize(string) bool
	Login(*browser.Browser, map[string]string) error
	Get(*browser.Browser, string) (string, *html.Node, error)
	Extract(string, *html.Node) (*Post, error)
}

type StrategyList struct {
	strats map[string]Strategy
	fallback   Strategy
}

func (sl *StrategyList) FindMatch(link string) Strategy {
	for _, s := range sl.strats {
		if s.Recognize(link) {
			return s
		}
	}
	return sl.fallback
}