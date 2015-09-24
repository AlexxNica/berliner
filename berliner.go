package berliner

import (
	"fmt"
	"sync"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
	"github.com/s3ththompson/berliner/log"
)

type Berliner struct {
	stream    stream
	renderers []func([]content.Post)
	errors	*errorWriter
}

func New() Berliner {
	w := newErrorWriter(1000)
	return Berliner{
		errors: w,
	}
}

func (b *Berliner) Go() {
	go func() {
		log.SetWriter(b.errors)
		for entry := range b.errors.entries() {
			fmt.Println("Filter error: " + entry.Message)
		}
	}()
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
	log.ResetWriter()
	b.errors.close()
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
