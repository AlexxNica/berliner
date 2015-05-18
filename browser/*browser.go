package browser

import (
	"errors"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
)

type Browser struct {
	bow	*browser.Browser
	sl *StrategyList
}

func New(credentials map[string]map[string]string) (*Browser, error) {
	b := &Browser {
		bow: surf.NewBrowser(),
		sl: &StrategyList{
			strats: map[string]Strategy{
				"new-york-times": &NewYorkTimes{},
				"new-yorker": &NewYorker{},
			},
			fallback: &Default{},
		},
	}
	b.bow.AddRequestHeader("Accept", "text/html")
	b.bow.AddRequestHeader("Accept-Charset", "utf8")
	for slug, creds := range credentials {
		s, ok := b.sl.strats[slug]
		if !ok {
			return nil, errors.New("unrecognized credential.")
		}
		err := s.Login(b.bow, creds)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

func (b *Browser) Browse(link string) (post *Post, err error) {
	s := b.sl.FindMatch(link)
	permalink, page, err := s.Get(b.bow, link)
	if err != nil {
		return nil, err
	}
	post, err = s.Extract(permalink, page)
	return
}
