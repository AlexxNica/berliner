package extractors

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
)

type Post struct {
	Title   string
	Content string
	Link    string
	Image   string
}

func (p *Post) String() string {
	return p.Title + "\t" + p.Content + "\t" + p.Link + "\t" + p.Image
}

type Extractor interface {
	Recognize(string) bool
	SetLink(string)
	Get() (*html.Node, error)
	Extract(*html.Node) (*Post, error)
	// Browser() browser of some sort
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

func New(link string) Extractor {
	m := &ExtractorList{
		extractors: []Extractor{},
		fallback:   &Default{},
	}
	e := m.FindMatch(link)
	e.SetLink(link)
	return e
}
