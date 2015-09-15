package content

import (
	"time"
)

type Post struct {
	Title     string
	Permalink string
	Body      string
	// Images    []Image // TODO
	// Videos    []Video
	Date     time.Time
	Authors  []string
	Tags     []string
	Source   string // TODO: rename source: we use source to mean feed aggregator AND individual source
	Via      string
	Language string
}

func MergePosts(p1, p2 Post) Post { // TODO: fix this shit
	if p2.Title != "" {
		p1.Title = p2.Title
	}
	if p2.Permalink != "" {
		p1.Permalink = p2.Permalink
	}
	if p2.Body != "" {
		p1.Body = p2.Body
	}
	if !p2.Date.IsZero() {
		p1.Date = p2.Date
	}
	if len(p2.Authors) != 0 {
		p1.Authors = p2.Authors
	}
	if len(p2.Tags) != 0 {
		p1.Tags = p2.Tags
	}
	if p2.Source != "" {
		p1.Source = p2.Source
	}
	if p2.Via != "" {
		p1.Via = p2.Via
	}
	if p2.Language != "" {
		p1.Language = p2.Language
	}
	return p1
}
