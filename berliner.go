package berliner

import (
	"sync"
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

type Berliner struct {
	options Options
	stream    stream
	renderers []renderer
}

type Options struct {
	Cadence time.Duration
	Refetch	bool
	Debug 	bool
}

const (
	Daily = 24*time.Hour
	Weekly = 7*Daily
	Forever = 0
)

func New(args ...Options) Berliner {
	options := NewOptions()
	if len(args) > 0 {
		options = args[0]
	}
	return Berliner{
		options: options,
	}
}

func NewOptions() Options {
	return Options{
		Cadence: Daily,
	}
}

func (b *Berliner) posts() <-chan content.Post {
	return b.stream.posts(scrape.NewClient(), b.options.Cadence)
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
	f func(*scrape.Client, time.Duration) <-chan content.Post
}

func (s source) posts(c *scrape.Client, d time.Duration) <-chan content.Post {
	return s.f(c, d)
}

func (b *Berliner) Source(name string, f func(*scrape.Client, time.Duration) <-chan content.Post) *stream {
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
