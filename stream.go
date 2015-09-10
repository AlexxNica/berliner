package berliner

import (
	"sync"
	. "github.com/s3ththompson/berliner/post"
)

type streamer interface {
	posts() <-chan Post
}

type source struct {
	f func() <-chan Post
}

func (s *source) posts() <-chan Post {
	return s.f()
}

type stream struct {
	children []streamer
	filters []func(<-chan Post) <-chan Post
}

func (s *stream) posts() <-chan Post {
	agg := make(chan Post)
	var wg sync.WaitGroup
	go func() {
		for _, child := range s.children {
			wg.Add(1)
			go func(child streamer) {
				for post := range child.posts() {
					agg <- post
				}
				wg.Done()
			}(child)
		}
		wg.Wait()
		close(agg)
	}()
	var out <-chan Post
	out = agg
	for _, filter := range s.filters {
		out = filter(out)
	}
	return out
}

func (s *stream) Filter(f func(<-chan Post) <-chan Post) {
	s.addFilter(f)
}

func (s *stream) addFilter(f func(<-chan Post) <-chan Post) {
	s.filters = append(s.filters, f)
}

func (s *stream) addSource(f func() <-chan Post) *stream {
	child := wrapSource(f)
	s.children = append(s.children, child)
	return child
}

func wrapSource(f func() <-chan Post) *stream {
	return &stream{
		children: []streamer{
			&source{
				f: f,
			},
		},
	}
}