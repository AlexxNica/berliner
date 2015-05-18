package main

import (
	"fmt"
	"net/url"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/SlyMarbo/rss"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/s3ththompson/berliner/extractors"
)

func Fetch(cmd *cobra.Command, args []string) {
	posts := make(chan *extractors.Post)
	p := &Pipe{
		workers: 10,
		do:      fetch,
		in:      readLines(),
		out:     posts,
	}
	err := p.pipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(extractors.PostHeader)
	for post := range posts {
		fmt.Println(post)
	}
}

func fetch(feed string, out chan<- *extractors.Post) {
	f, err := rss.Fetch(feed)
	if err != nil {
		return
	}
	for _, item := range f.Items {
		link := item.ID
		_, err := url.Parse(link)
		if err != nil {
			continue
		}
		out <- &extractors.Post{
			Permalink: link,
			Title:     item.Title,
			Date:      item.Date,
		}
	}
}
