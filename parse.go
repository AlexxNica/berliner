package main

import (
	"fmt"

	"github.com/s3ththompson/berliner/extractors"
	// "github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func Parse(cmd *cobra.Command, args []string) {
	posts := make(chan *extractors.Post)
	p := &Pipe{
		workers: 20,
		do:      parse,
		in:      readLines(),
		out:     posts,
	}
	err := p.pipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	for post := range posts {
		fmt.Println(post)
	}
}

func parse(link string, out chan *extractors.Post) {
	e := extractors.New(link)
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
