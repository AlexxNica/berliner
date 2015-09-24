package berliner

import (
	"sync"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

type Berliner struct {
	stream    stream
	renderers []renderer
}

func New() Berliner {
	return Berliner{}
}

func (b *Berliner) posts() <-chan content.Post {
	return b.stream.posts(scrape.NewClient())
}

func (b *Berliner) render(posts []content.Post) {
	var wg sync.WaitGroup
	for _, r := range b.renderers {
		wg.Add(1)
		go func(r renderer) {
			defer wg.Done()
			r.f(posts)
		}(r)
	}
	wg.Wait()
}

func (b *Berliner) Go() {
	posts := gather(b.posts())
	b.render(posts)
}

type source struct {
	name string
	f func(*scrape.Client) <-chan content.Post
}

func (s source) posts(c *scrape.Client) <-chan content.Post {
	return s.f(c)
}

func (b *Berliner) Source(name string, f func(*scrape.Client) <-chan content.Post) *stream {
	return b.addSource(source{
		name: name,
		f: f,
	})
}

func (b *Berliner) addSource(source source) *stream {
	return b.stream.addSource(source)
}

type filter struct {
	name string
	f func(<-chan content.Post) <-chan content.Post
}

func (b *Berliner) Filter(name string, f func(<-chan content.Post) <-chan content.Post) {
	b.addFilter(filter{
		name: name,
		f: f,
	})
}

func (b *Berliner) addFilter(filter filter) {
	b.stream.addFilter(filter)
}

type renderer struct {
	name string
	f func([]content.Post)
}

func (b *Berliner) Renderer(name string, f func([]content.Post)) {
	b.addRenderer(renderer{
		name: name,
		f: f,
	})
}

func (b *Berliner) addRenderer(renderer renderer) {
	b.renderers = append(b.renderers, renderer)
}

func (b *Berliner) Posts() <-chan content.Post {
	return b.posts()
}
