package content

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/microcosm-cc/bluemonday"
)

type Post struct {
	Title     string
	Permalink string
	Body      string
	Images    []string
	// Videos    []Video
	Date     time.Time
	Authors  []string
	Tags     []string
	Origin   string
	Via      string
	Language string
	Points   int
}

func (p *Post) Sanitize() {
	sanitized := bluemonday.UGCPolicy().Sanitize(p.Body)
	p.Body = sanitized
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
	if len(p2.Images) != 0 {
		p1.Images = p2.Images
	}
	if p2.Origin != "" {
		p1.Origin = p2.Origin
	}
	if p2.Via != "" {
		p1.Via = p2.Via
	}
	if p2.Language != "" {
		p1.Language = p2.Language
	}
	return p1
}

// Returns the word count of the post.
func (p *Post) Wordcount() int {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(p.Body))
	if err != nil {
		fmt.Println(os.Stderr, "Error: Couldn't compute word count.")
		return 0
	}
	text := doc.Find("p").Text() // only count text in paragraph tags
	return len(strings.Fields(text))
}
