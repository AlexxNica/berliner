package main

import (
	"fmt"
	"os"

	"github.com/s3ththompson/berliner/browser"
	// "github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
)

func Parse(cmd *cobra.Command, args []string) {
	posts := make(chan *browser.Post)
	p := &Pipe{
		workers: 20,
		do:      parse,
		in:      readPosts(),
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

func readPosts() <-chan *browser.Post {
	out := make(chan *browser.Post)
	go func() {
		defer close(out)
		posts, err := ReadPosts(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		for _, post := range posts {
			out <- post
		}
	}()
	return out
}

func parse(p *browser.Post, out chan<- *browser.Post) {
	b, err := browser.New(map[string]map[string]string{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	post, err := b.Parse(p.Permalink)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	out <- post
}
