package browser

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
)

type Strategy interface {
	slug() string
	recognize(string) bool
	login(*browser.Browser, map[string]string) error
	get(*browser.Browser, string) (string, *html.Node, error)
	extract(string, *html.Node) (*Post, error)
}

type Strategies struct {
	strats   map[string]Strategy
	fallback Strategy
}

func (s *Strategies) bySlug(slug string) (Strategy, bool) {
	strat, ok := s.strats[slug]
	return strat, ok
}

func (s *Strategies) byLink(link string) Strategy {
	for _, strat := range s.strats {
		if strat.recognize(link) {
			return strat
		}
	}
	return s.fallback
}

var strategies Strategies = Strategies{
	fallback: &fallback{},
}

func register(strat Strategy) {
	slug := strat.slug()
	if _, dup := strategies.strats[slug]; dup {
		return
	}
	strategies.strats[slug] = strat
}
