package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
	"net/http"
	"sync"
)

func pipeRead(urls <-chan string) <-chan *html.Node {
	out := make(chan *html.Node)
	go func() {
		var wg sync.WaitGroup
		for url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				read(url, out)
			}(url)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func read(url string, out chan *html.Node) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Read error: %s\n", url)
		return
	}
	r, err := charset.NewReader(resp.Body, resp.Header.Get("content-type"))
	if err != nil {
		fmt.Printf("UTF8 error: %s\n", url)
		return
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Printf("Parse error: %s\n", url)
		return
	}
	out <- doc
}
