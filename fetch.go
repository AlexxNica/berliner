package main

import (
	"fmt"
	"net/url"

	"github.com/SlyMarbo/rss"
	"github.com/spf13/cobra"
)

func Fetch(cmd *cobra.Command, args []string) {
	links := make(chan string)
	p := &Pipe{
		workers: 10,
		do:      fetch,
		in:      readLines(),
		out:     links,
	}
	err := p.pipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	for link := range links {
		fmt.Println(link)
	}
}

func fetch(feed string, out chan<- string) {
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
		out <- link
	}
}
