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

func (b *Browser) massLogin(credentials map[string]map[string]string) error {
	for slug, creds := range credentials {
		s, ok := b.sl.strats[slug]
		if !ok {
			return errors.New("unrecognized credential.")
		}
		err := s.login(b.bow, creds)
		if err != nil {
			return err
		}
	}
	return nil
}

func New(credentials map[string]map[string]string) (b *Browser, err error) {
	b = &Browser {
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
	err = b.massLogin(credentials)
	return
}

func (b *Browser) Parse(link string) (post *Post, err error) {
	s := b.sl.findMatch(link)
	permalink, page, err := s.get(b.bow, link)
	if err != nil {
		return nil, err
	}
	post, err = s.extract(permalink, page)
	if err != nil {
		return nil, err
	}
	err = post.sanitize()
	return post, err
}
