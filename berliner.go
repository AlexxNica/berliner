package berliner

import (
	"sync"

	. "github.com/s3ththompson/berliner/post"
)

type Berliner struct {
	stream stream
	renderers []func([]Post)
}

func New() Berliner {
	return Berliner{}
}

func (b *Berliner) Go() {
	posts := []Post{}
	for post := range b.stream.posts() {
		posts = append(posts, post)
	}
	var wg sync.WaitGroup
	for _, renderer := range b.renderers {
		wg.Add(1)
		go func(renderer func([]Post)) {
			defer wg.Done()
			renderer(posts)
		}(renderer)
	}
	wg.Wait()
}

func (b *Berliner) Renderer(f func([]Post)) {
	b.renderers = append(b.renderers, f)
}

func (b *Berliner) Filter(f func(<-chan Post) <-chan Post) {
	b.stream.addFilter(f)
}

func (b *Berliner) Source(f func() <-chan Post) *stream {
	return b.stream.addSource(f)
}

func (b *Berliner) Posts() <-chan Post {
	return b.stream.posts()
}