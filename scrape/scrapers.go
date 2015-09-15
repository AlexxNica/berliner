package scrape

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/s3ththompson/berliner/content"
)

type scraper interface {
	recognize(string) bool
	scrape(*html.Node) (content.Post, error)
}

type lookup struct {
	list     []scraper
	fallback scraper
}

func (l *lookup) byURL(url string) scraper {
	for _, s := range l.list {
		if s.recognize(url) {
			return s
		}
	}
	return l.fallback
}

var scrapers lookup = lookup{
	list:     []scraper{},
	fallback: &fallback{},
}

func register(s scraper) {
	scrapers.list = append(scrapers.list, s)
}
