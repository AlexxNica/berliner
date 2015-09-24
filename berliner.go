package berliner

import (
	"errors"
	"sync"
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

type Berliner struct {
	cache		*Cache
	options   	Options
	stream    	stream
	renderers 	[]renderer
}

type Options struct {
	Cache 		bool
	CacheFile	string
	Cadence 	time.Duration
	Debug   	bool
}

const (
	Daily   = 24 * time.Hour
	Weekly  = 7 * Daily
	Forever = 0
)

// TODO: document that berliners themselves are not threadsafe
func New(args ...Options) (Berliner, error) {
	options := NewOptions()
	if len(args) > 0 {
		options = args[0]
	}
	b := Berliner{
		options: options,
	}
	if b.options.Cache {
		if b.options.CacheFile == "" {
			return Berliner{}, errors.New("no cache file specified")
		}
		cache, err := NewCache(b.options.CacheFile)
		if err != nil {
			return Berliner{}, err
		}
		b.cache = cache
	}
	return b, nil
}

func NewOptions() Options {
	return Options{
		Cadence: Daily,
	}
}

func (b *Berliner) posts() <-chan content.Post {
	var client scrape.Client
	if b.options.Cache {
		client = NewCachedClient(b.cache)
	} else {
		client = scrape.NewClient()
	}
	return b.stream.posts(client, b.options.Cadence)
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

func (b *Berliner) CleanUp() error {
	return b.cache.Close()
}

type source struct {
	name string
	f    func(scrape.Client, time.Duration) <-chan content.Post
}

func (s source) posts(c scrape.Client, d time.Duration) <-chan content.Post {
	return s.f(c, d)
}

func (b *Berliner) Source(name string, f func(scrape.Client, time.Duration) <-chan content.Post) *stream {
	return b.addSource(source{
		name: name,
		f:    f,
	})
}

func (b *Berliner) addSource(source source) *stream {
	return b.stream.addSource(source)
}

type filter struct {
	name string
	f    func(<-chan content.Post) <-chan content.Post
}

func (b *Berliner) Filter(name string, f func(<-chan content.Post) <-chan content.Post) {
	b.addFilter(filter{
		name: name,
		f:    f,
	})
}

func (b *Berliner) addFilter(filter filter) {
	b.stream.addFilter(filter)
}

type renderer struct {
	name string
	f    func([]content.Post)
}

func (b *Berliner) Renderer(name string, f func([]content.Post)) {
	b.addRenderer(renderer{
		name: name,
		f:    f,
	})
}

func (b *Berliner) addRenderer(renderer renderer) {
	b.renderers = append(b.renderers, renderer)
}

func (b *Berliner) Posts() <-chan content.Post {
	return b.posts()
}
