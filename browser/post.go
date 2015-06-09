package browser

import (
	"time"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/microcosm-cc/bluemonday"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/rubenfonseca/fastimage"
)

type Image struct {
	URL    string
	Alt    string
	Width  uint32
	Height uint32
}

func newImage(url, alt string) Image {
	i := Image{
		URL: url,
		Alt: alt,
	}
	_, size, err := fastimage.DetectImageType(i.URL)
	if err == nil {
		i.Width = size.Width
		i.Height = size.Height
	}
	return i
}

type Movie struct {
	URL string
	Alt string
}

func newMovie(url, alt string) Movie {
	m := Movie{
		URL: url,
		Alt: alt,
	}
	return m
}

type Post struct {
	Title     string
	Permalink string
	Content   string
	Images    []Image
	Movies    []Movie
	Date      time.Time
	Authors   []string
	Tags      []string
	Source    string
	Language  string
}

func (p *Post) sanitize() {
	sanitized := bluemonday.UGCPolicy().Sanitize(p.Content)
	p.Content = sanitized
}

func (p Post) validate() bool {
	return true
}

func (p *Post) addImage(url string, alt string) {
	p.Images = append(p.Images, newImage(url, alt))
}

func (p *Post) addMovie(url string, alt string) {
	p.Movies = append(p.Movies, newMovie(url, alt))
}

func (p Post) String() string {
	return p.Source + ": " + p.Title
}
