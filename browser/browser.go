package browser

import (
	"errors"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/headzoo/surf/browser"
)

type Browser struct {
	Bow *browser.Browser
	sl  *StrategyList
}

func (b *Browser) Init() {
	b.Bow = surf.NewBrowser()
	b.sl = &StrategyList{
		strats: map[string]Strategy{
			"new-york-times": &NewYorkTimes{},
			"new-yorker":     &NewYorker{},
		},
		fallback: &Default{},
	}
	b.Bow.AddRequestHeader("Accept", "text/html")
	b.Bow.AddRequestHeader("Accept-Charset", "utf8")
}

func (b *Browser) Login(credentials map[string]map[string]string) error {
	for slug, creds := range credentials {
		s, ok := b.sl.strats[slug]
		if !ok {
			return errors.New("unrecognized credential.")
		}
		err := s.login(b.Bow, creds)
		if err != nil {
			return err
		}
	}
	return nil
}

func New(credentials map[string]map[string]string) (b *Browser, err error) {
	b = &Browser{}
	b.Init()
	err = b.Login(credentials)
	return
}

func (b *Browser) Parse(link string) (post *Post, err error) {
	s := b.sl.findMatch(link)
	permalink, page, err := s.get(b.Bow, link)
	if err != nil {
		return nil, err
	}
	post, err = s.extract(permalink, page)
	if err != nil {
		return nil, err
	}
	return post, err
}
