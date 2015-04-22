package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/SlyMarbo/rss"
	"github.com/spf13/cobra"
)

func Fetch(cmd *cobra.Command, args []string) {
	feeds := readLines()
	links := fetch(feeds)
	for link := range links {
		fmt.Fprintln(os.Stdout, link)
	}
}

func fetch(feeds <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		var wg sync.WaitGroup

		for feed := range feeds {
			wg.Add(1)
			go func(feed string) {
				defer wg.Done()
				innerFetch(feed, out)
			}(feed)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func innerFetch(feed string, out chan string) {
	f, err := rss.Fetch(feed)
	if err != nil {
		return
	}
	for _, item := range f.Items {
		out <- item.ID
	}
}
