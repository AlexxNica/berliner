package berliner

import (
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
	"sync"
)

type streamer interface {
	posts(*scrape.Client) <-chan content.Post
}

type source struct {
	f func(*scrape.Client) <-chan content.Post
}

func (s *source) posts(c *scrape.Client) <-chan content.Post {
	return s.f(c)
}

type stream struct {
	children []streamer
	filters  []func(<-chan content.Post) <-chan content.Post
}

func (s *stream) posts(c *scrape.Client) <-chan content.Post {
	agg := make(chan content.Post)
	var wg sync.WaitGroup
	go func() {
		for _, child := range s.children {
			wg.Add(1)
			go func(child streamer) {
				for post := range child.posts(c) {
					agg <- post
				}
				wg.Done()
			}(child)
		}
		wg.Wait()
		close(agg)
	}()
	var out <-chan content.Post
	out = agg
	for _, filter := range s.filters {
		out = filter(out)
	}
	return out
}

func (s *stream) Filter(f func(<-chan content.Post) <-chan content.Post) {
	s.addFilter(f)
}

func (s *stream) addFilter(f func(<-chan content.Post) <-chan content.Post) {
	s.filters = append(s.filters, f)
}

func (s *stream) addSource(f func(*scrape.Client) <-chan content.Post) *stream {
	child := wrapSource(f)
	s.children = append(s.children, child)
	return child
}

func wrapSource(f func(*scrape.Client) <-chan content.Post) *stream {
	return &stream{
		children: []streamer{
			&source{
				f: f,
			},
		},
	}
}
