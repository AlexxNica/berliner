package post

import (
	"time"
)

type Post struct {
	Title     string
	Permalink string
	Content   string
	// Images    []Image
	// Videos    []Video
	Date      time.Time
	Authors   []string
	Tags      []string
	Source    string
	Via       string
	Language  string
}