package extractors

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
)

type Extractor interface {
	Recognize(string) bool
	SetPost(*Post)
	Get() (*html.Node, error)
	Extract(*html.Node) (*Post, error)
}

type ExtractorList struct {
	extractors []Extractor
	fallback   Extractor
}

func (m *ExtractorList) FindMatch(link string) Extractor {
	for _, e := range m.extractors {
		if e.Recognize(link) {
			return e
		}
	}
	return m.fallback
}

func New(post *Post) Extractor {
	m := &ExtractorList{
		extractors: []Extractor{},
		fallback:   &Default{},
	}
	e := m.FindMatch(post.Permalink)
	e.SetPost(post)
	return e
}
