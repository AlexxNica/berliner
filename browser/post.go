package browser

import (
	"strings"
	"time"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"github.com/rubenfonseca/fastimage"
)

type Image struct {
	URL string
	Alt string
	Width uint32
	Height uint32
}

func (i *Image) fastImage() {
	_, size, err := fastimage.DetectImageType(i.URL)
	if (err != nil) {
		return
	}
	i.Width = size.Width
	i.Height = size.Height
}

func (i Image) String() string {
	return i.URL
}

type Movie struct {
	URL string
	Alt string
}

func (m Movie) String() string {
	return m.URL
}

type Post struct {
	title     string
	permalink string
	content   *html.Node
	images    []Image
	movies    []Movie
	date      time.Time
	authors   []string
	tags      []string
	source    string
	language  string
}

func (p *Post) addImage(url string, alt string) {
	p.images = append(p.images, Image{
			URL: url,
			Alt: alt,
		})
}

func (p *Post) addMovie(url string, alt string) {
	p.movies = append(p.movies, Movie{
			URL: url,
			Alt: alt,
		})
}

func (p *Post) setContent(s string) {
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		doc = nil
	}
	p.content = doc
}

// TODO: actual sanitizing
func (p Post) sanitized() *html.Node {
	return p.content
}

func (p Post) String() string {
	return p.Source() + ": " + p.Title()
}

func (p Post) Title() string {
	return trim(squeeze(p.title))
}

func (p Post) Permalink() string {
	return trim(squeeze(p.permalink))
}

func (p Post) Content() string {
	// TODO: remove squeeze when client-side code is fixed
	return squeeze(render(p.sanitized()))
}

func (p Post) ContentNode() *html.Node {
	return p.sanitized()
}

func (p Post) Images() []Image {
	var out []Image
	for _, image := range p.images {
		image.fastImage()
		out = append(out, image)
	}
	return out
}

func (p Post) Movies() []Movie {
	return p.movies
}

func (p Post) Date() time.Time {
	return p.date
}

func (p Post) Authors() []string {
	var out []string
	for _, author := range p.authors {
		out = append(out, trim(squeeze(author)))
	}
	return out
}

func (p Post) Tags() []string {
	var out []string
	for _, tag := range p.tags {
		out = append(out, trim(squeeze(tag)))
	}
	return out
}

func (p Post) Source() string {
	return p.source
}

// TODO: enum?
func (p Post) Langauge() string {
	return p.language
}