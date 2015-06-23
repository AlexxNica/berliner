package browser

import (
	"regexp"
	"time"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/microcosm-cc/bluemonday"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/rubenfonseca/fastimage"
)

var rxURL = regexp.MustCompile(`^((ftp|http|https):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|((www\.)?)?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?_?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?)|localhost)(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`)

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

type Video struct {
	URL string
	Alt string
}

func newVideo(url, alt string) Video {
	m := Video{
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
	Videos    []Video
	Date      time.Time
	Authors   []string
	Tags      []string
	Source    string
	Via       string
	Language  string
}

func (p *Post) sanitize() {
	sanitized := bluemonday.UGCPolicy().Sanitize(p.Content)
	p.Content = sanitized
}

func (p Post) validate() bool {
	for _, image := range p.Images {
		if s := image.URL; s == "" || !rxURL.MatchString(s) {
			return false
		}
	}
	for _, video := range p.Videos {
		if s := video.URL; s == "" || !rxURL.MatchString(s) {
			return false
		}
	}
	if !rxURL.MatchString(p.Permalink) {
		return false
	}
	valid := (p.Title != "" && p.Permalink != "" && p.Content != "" && p.Source != "")
	return valid
}

func (p *Post) addImage(url string, alt string) {
	p.Images = append(p.Images, newImage(url, alt))
}

func (p *Post) addVideo(url string, alt string) {
	p.Videos = append(p.Videos, newVideo(url, alt))
}
