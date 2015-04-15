package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// conf := readConf()
	input := []string{
		"http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
		"http://www.vox.com/rss/index.xml",
	}

	feeds := startPipe(input)
	urls := pipeFetch(feeds)
	htmls := pipeRead(urls)
	// u := filterURLs(urls)
	// htmls := read(urls)
	// posts := parse(htmls)
	// p := filterPosts(posts)
	// out := collect(posts)

	// render(out)

	for html := range htmls {
		doc := goquery.NewDocumentFromNode(html)
		fmt.Printf("Title: %s\n", doc.Find("head title").Text())
	}
}

// temporary function, but equivalent to gen()
// in this example https://blog.golang.org/pipelines
func startPipe(feeds []string) <-chan string {
	// oh yeah, unbuffered channels block receivers
	// until data is available on the channel
	// AND they block senders until receivers are ready
	// to receive downstream.  thus the senders here
	// need to be async if we want to channel to be unbuffered
	out := make(chan string)
	go func() {
		for _, f := range feeds {
			out <- f
		}
		// in all of our examples, closing the pipe means
		// telling the downstream functions that there's no
		// more work to be done
		close(out)
	}()
	return out
}
