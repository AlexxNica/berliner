package main

import (
	"fmt"
)

func main() {
	// conf := readConf()
	input := []string{
		"http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
		"http://www.vox.com/rss/index.xml",
		}

	feeds := startPipe(input)
	urls := pipeFetch(feeds)
	// htmls := pipeRead(urls)
	// u := filterURLs(urls)
	// htmls := read(urls)
	// posts := parse(htmls)
	// p := filterPosts(posts)
	// out := collect(posts)

	// render(out)

	for url := range urls {
		fmt.Printf("%s\n", url)
	}
}

// temporary function, but equivalent to gen()
// in this example https://blog.golang.org/pipelines
func startPipe(feeds []string) <-chan string {
	// adding a length parameter to a channel makes
	// it buffered (rather than unbuffered)
	// I can't remember why this needs to be here :(
	// I'll look it up
	out := make(chan string, len(feeds))
	for _, f := range feeds {
		out <- f
	}
	// in all of our examples, closing the pipe means
	// telling the downstream functions that there's no
	// more work to be done
	close(out)
	return out
}