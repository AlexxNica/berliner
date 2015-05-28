package browser

import (
	"time"
)

type Post struct {
	Title     string
	Permalink string
	Content   string
	Images    []string
	Movies    []string
	Date      time.Time
	Authors   []string
	Tags      []string
	Source    string
	Language  string
}

func (p *Post) sanitize() (err error) {
	return
}