package berliner

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/log"
	"github.com/s3ththompson/berliner/scrape"

	"github.com/fatih/color"
)

type Berliner struct {
	cache     *Cache
	options   Options
	stream    *stream
	renderers []*renderer
}

type Options struct {
	Cache     bool
	CacheFile string
	Cadence   time.Duration
	Debug     bool
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
		stream:  &stream{},
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

func (b *Berliner) walk() {
	walkStream(b.stream, 0)
	for _, renderer := range b.renderers {
		renderer.init(0)
	}
}

// TODO: make less dangerous???
func walkStream(s *stream, indent int) {
	// check if streamer is a wrapped source
	if len(s.children) == 1 && reflect.TypeOf(s.children[0]).String() == "*berliner.source" {
		s.children[0].(*source).init(indent)
	} else {
		for _, child := range s.children {
			walkStream(child.(*stream), indent+1)
		}
	}
	for _, filter := range s.filters {
		filter.init(indent)
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
		go func(r *renderer) {
			defer wg.Done()
			r.f(posts)
		}(r)
	}
	wg.Wait()
}

func (b *Berliner) Go() {
	b.walk() // for now, just set up log entries and print initial skeleton
	posts := gather(b.posts())
	b.render(posts)
}

func (b *Berliner) CleanUp() error {
	if b.options.Cache {
		return b.cache.Close()
	}
	return nil
}

type source struct {
	name  string
	f     func(scrape.Client, time.Duration) <-chan content.Post
	entry *log.Entry
}

func (s *source) init(indent int) {
	s.entry = log.Println(sourceMessage(indent, s.name, -1, false))
	oldf := s.f
	s.f = func(c scrape.Client, d time.Duration) <-chan content.Post {
		seen := 0
		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range oldf(c, d) {
				seen++
				s.entry.Updateln(sourceMessage(indent, s.name, seen, false))
				out <- post
			}
			s.entry.Updateln(sourceMessage(indent, s.name, seen, true))
		}()
		return out
	}
}

func (s *source) posts(c scrape.Client, d time.Duration) <-chan content.Post {
	return s.f(c, d)
}

func (b *Berliner) Source(name string, f func(scrape.Client, time.Duration) <-chan content.Post) *stream {
	return b.addSource(&source{
		name: name,
		f:    f,
	})
}

func (b *Berliner) addSource(source *source) *stream {
	return b.stream.addSource(source)
}

type filter struct {
	name  string
	f     func(<-chan content.Post) <-chan content.Post
	entry *log.Entry
}

func (f *filter) init(indent int) {
	f.entry = log.Println(filterMessage(indent, f.name, -1, -1, false))
	oldf := f.f
	f.f = func(posts <-chan content.Post) <-chan content.Post {
		seenIn := 0
		seenOut := 0
		in := make(chan content.Post)
		go func() {
			defer close(in)
			for post := range posts {
				seenIn++
				f.entry.Updateln(filterMessage(indent, f.name, seenIn, seenOut, false))
				in <- post
			}
		}()
		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range oldf(in) {
				seenOut++
				f.entry.Updateln(filterMessage(indent, f.name, seenIn, seenOut, false))
				out <- post
			}
			f.entry.Updateln(filterMessage(indent, f.name, seenIn, seenOut, true))
		}()
		return out
	}
}

func (b *Berliner) Filter(name string, f func(<-chan content.Post) <-chan content.Post) {
	b.addFilter(&filter{
		name: name,
		f:    f,
	})
}

func (b *Berliner) addFilter(filter *filter) {
	b.stream.addFilter(filter)
}

type renderer struct {
	name  string
	f     func([]content.Post)
	entry *log.Entry
}

func (r *renderer) init(indent int) {
	r.entry = log.Println(rendererMessage(indent, r.name, false))
	oldf := r.f
	r.f = func(posts []content.Post) {
		oldf(posts)
		r.entry.Updateln(rendererMessage(indent, r.name, true))
	}
}

func (b *Berliner) Renderer(name string, f func([]content.Post)) {
	b.addRenderer(&renderer{
		name: name,
		f:    f,
	})
}

func (b *Berliner) addRenderer(renderer *renderer) {
	b.renderers = append(b.renderers, renderer)
}

func (b *Berliner) Posts() <-chan content.Post {
	return b.posts()
}

func sourceMessage(indent int, name string, out int, done bool) string {
	var o string
	if out < 0 {
		o = "-"
	} else {
		o = strconv.Itoa(out)
	}
	s := fmt.Sprintf("Fetching %s posts: %s out...", color.WhiteString(name), color.GreenString(o))
	if done {
		s += " done"
	}
	s = "├─ " + s
	return s
}

func filterMessage(indent int, name string, in, out int, done bool) string {
	var i string
	if in < 0 {
		i = "-"
	} else {
		i = strconv.Itoa(in)
	}
	var o string
	if out < 0 {
		o = "-"
	} else {
		o = strconv.Itoa(out)
	}
	s := fmt.Sprintf("Running %s filter: %s in, %s out...", color.WhiteString(name), color.GreenString(i), color.GreenString(o))
	if done {
		s += " done"
	}
	if indent > 0 {
		s = "│ ├─ " + s
	} else {
		s = "├─ " + s
	}
	return s
}

func rendererMessage(indent int, name string, done bool) string {
	s := fmt.Sprintf("Rendering %s...", color.WhiteString(name))
	if done {
		s += " done"
	}
	s = "├─ " + s
	return s
}
