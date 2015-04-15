package main

import (
	"github.com/SlyMarbo/rss"
	"sync"
)

func pipeFetch(urls <-chan string) <-chan string {
	out := make(chan string)
	// going async here isn't really the big speed benefit
	// but it's still a healthy idiom
	go func() {
		// this comes straight from that blog post.
		// the difficult part is knowing when to close
		// the out chan
		// fetching is asynchronous, so we can't just close
		// it after we start the fetches... we need to close
		// it after the last fetch finishes
		// the wait group waits for a group of work to finish
		var wg sync.WaitGroup
		// it's worth looking up the behavior of range on a
		// channel--it's different for range on a slice/array
		for url := range urls {
			// tell the wait group to wait for this url's done
			// signal
			wg.Add(1)
			// this is where the real async speed benefits come
			go func(url string) {
				// defer is a little more idiomatic, but functionally equivalent
				// to putting wg.Done() after our synchronous fetch function
				defer wg.Done()
				fetch(url, out)
				// the reason we refer to out using a closure, but pass
				// in url explicitly is one of few go gotchas
				// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
			}(url)
		}
		// we arrive here immediately after the anonymous
		// functions w/ fetch are kicked off, but we need to
		// wait until all the fetching actually occurs
		wg.Wait()
		close(out)
	}()
	return out
}

func fetch(url string, out chan string) {
	feed, err := rss.Fetch(url)
	// explicit error catching like this is verbose
	// but very idiomatic
	// you might come to like how safe it seems to always
	// handle exceptions, at the point that they first arise,
	// rather than waiting for them to bubble up elsewhere
	// here, if the fetch fails, we make the conscious decision
	// to just skip sending the post url downstream
	if err != nil {
		return
	}
	for _, item := range feed.Items {
		// this is currently a problem:
		// looks like some feeds put permalink in
		// item.Link, and some in item.ID
		out <- item.ID
	}
}
