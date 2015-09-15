package berliner

import (
	"sync"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

type Berliner struct {
	stream    stream
	renderers []func([]content.Post)
}

func New() Berliner {
	return Berliner{}
}

func (b *Berliner) Go() {
	posts := []content.Post{}
	for post := range b.stream.posts(scrape.NewClient()) {
		posts = append(posts, post)
	}
	var wg sync.WaitGroup
	for _, renderer := range b.renderers {
		wg.Add(1)
		go func(renderer func([]content.Post)) {
			defer wg.Done()
			renderer(posts)
		}(renderer)
	}
	wg.Wait()
}

func (b *Berliner) Renderer(f func([]content.Post)) {
	b.renderers = append(b.renderers, f)
}

func (b *Berliner) Filter(f func(<-chan content.Post) <-chan content.Post) {
	b.stream.addFilter(f)
}

func (b *Berliner) Source(f func(*scrape.Client) <-chan content.Post) *stream {
	return b.stream.addSource(f)
}

func (b *Berliner) Posts() <-chan content.Post {
	return b.stream.posts(scrape.NewClient())
}
