package main

import (
	"encoding/csv"
	"io"
	"strings"
	"time"

	"github.com/s3ththompson/berliner/browser"
)

func WritePosts(p []*browser.Post, w io.Writer) error {
	writer := csv.NewWriter(w)
	writer.Comma = '\t'
	for _, post := range p {
		err := writer.Write(postToSlice(post))
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func ReadPosts(r io.Reader) ([]*browser.Post, error) {
	reader := csv.NewReader(r)
	reader.Comma = '\t'
	posts, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	out := []*browser.Post{}
	for _, post := range posts {
		out = append(out, sliceToPost(post))
	}
	return out, nil
}

func imagesToSlice(images []browser.Image) []string {
	out := []string{}
	for _, image := range images {
		out = append(out, image.URL)
	}
	return out
}

func sliceToImages(s []string) []browser.Image {
	out := []browser.Image{}
	for _, url := range s {
		out = append(out, browser.Image{URL: url})
	}
	return out
}

func videosToSlice(videos []browser.Video) []string {
	out := []string{}
	for _, video := range videos {
		out = append(out, video.URL)
	}
	return out
}

func sliceToVideos(s []string) []browser.Video {
	out := []browser.Video{}
	for _, url := range s {
		out = append(out, browser.Video{URL: url})
	}
	return out
}

func postToSlice(p *browser.Post) []string {
	return []string{
		p.Title,
		p.Permalink,
		p.Content,
		strings.Join(imagesToSlice(p.Images), ","),
		strings.Join(videosToSlice(p.Videos), ","),
		p.Date.Format(time.RFC3339),
		strings.Join(p.Authors, ","),
		strings.Join(p.Tags, ","),
		p.Source,
		p.Via,
		p.Language,
	}
}

func sliceToPost(s []string) *browser.Post {
	date, err := time.Parse(time.RFC3339, s[5])
	if err != nil {
		date = time.Now()
	}
	return &browser.Post{
		Title:     s[0],
		Permalink: s[1],
		Content:   s[2],
		Images:    sliceToImages(strings.Split(s[3], ",")),
		Videos:    sliceToVideos(strings.Split(s[4], ",")),
		Date:      date,
		Authors:   strings.Split(s[6], ","),
		Tags:      strings.Split(s[7], ","),
		Source:    s[8],
		Via:       s[9],
		Language:  s[10],
	}
}
