package main

import (
	"fmt"
	"os"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/SlyMarbo/rss"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/s3ththompson/berliner/browser"
)

func Fetch(cmd *cobra.Command, args []string) {
	posts := make(chan *browser.Post)
	p := &Pipe{
		workers: 10,
		do:      fetch,
		in:      readLines(),
		out:     posts,
	}
	err := p.pipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	out := []*browser.Post{}
	for post := range posts {
		out = append(out, post)
	}
	err = WritePosts(out, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func fetch(feed string, out chan<- *browser.Post) {
	// ignore blank or commented out lines
	if len(feed) == 0 || feed[0] == '#' {
		return
	}

	f, err := rss.Fetch(feed)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, item := range f.Items {
		link := item.Link
		if link == "" {
			link = item.ID
		}
		out <- &browser.Post{
			Permalink: link,
			Title:     item.Title,
			Date:      item.Date,
		}
	}
}
