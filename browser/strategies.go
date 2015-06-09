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

type _strategies struct {
	strats   map[string]strategy
	fallback strategy
}

func (s *_strategies) bySlug(slug string) (strategy, bool) {
	strat, ok := s.strats[slug]
	return strat, ok
}

func (s *_strategies) byLink(link string) strategy {
	for _, strat := range s.strats {
		if strat.recognize(link) {
			return strat
		}
	}
	return s.fallback
}

var strategies _strategies = _strategies{
	strats:   make(map[string]strategy),
	fallback: &fallback{},
}

func register(strat strategy) {
	slug := strat.slug()
	if _, dup := strategies.strats[slug]; dup {
		return
	}
	strategies.strats[slug] = strat
}
