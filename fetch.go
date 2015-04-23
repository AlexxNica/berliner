package main

import (
	"fmt"
	"os"

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
		fmt.Fprintln(os.Stdout, link)
	}
}

func fetch(feed string, out chan<- string) {
	f, err := rss.Fetch(feed)
	if err != nil {
		return
	}
	for _, item := range f.Items {
		out <- item.ID
	}
}
