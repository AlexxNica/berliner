package main

import (
	"fmt"
	"github.com/s3ththompson/berliner/extractors"
	"os"
	// "github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
)

func Parse(cmd *cobra.Command, args []string) {
	posts := make(chan *extractors.Post)
	p := &Pipe{
		workers: 20,
		do:      parse,
		in:      readPosts(),
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

func readPosts() <-chan *extractors.Post {
	out := make(chan *extractors.Post)
	go func() {
		posts, err := extractors.ReadPosts(os.Stdin)
		if err == nil {
			for _, post := range posts {
				out <- post
			}
		}
		close(out)
	}()
	return out
}

func parse(p *extractors.Post, out chan<- *extractors.Post) {
	e := extractors.New(p)
	page, err := e.Get()
	if err != nil {
		return
	}
	post, err := e.Extract(page)
	if err != nil {
		return
	}
	out <- post
}
