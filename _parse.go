package main

import (
	"fmt"
	"github.com/s3ththompson/berliner/browser"
	"os"
	// "github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
)

var cmdParse = &cobra.Command{
	Use:   "_parse",
	Short: "Parse article permalinks",
	Long:  "Parse structured articles to stdout from permalinks from stdin",
}

func init() {
	cmdParse.Run = Parse
}


func Parse(cmd *cobra.Command, args []string) {
	posts := make(chan *Post)
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
	fmt.Println(PostHeader)
	for post := range posts {
		fmt.Println(post)
	}
}

func readPosts() <-chan *Post {
	out := make(chan *Post)
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

func parse(p *Post, out chan<- *Post) {
	b, err := browser.New(map[string]map[string]string{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	post, err := b.Parse(p.Permalink)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	mypost := Post(*post)
	out <- &mypost // cast *browser.Post to *Post
}
