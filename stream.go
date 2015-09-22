package berliner

import (
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
	"sync"
)

// a stream is a set of "things that can emit posts" + filters to apply to those posts.
// "things that can emit posts" are usually sources, but they can also be child streams,
// thus enabling the ability to combine multiple sets of sources + filters in a proper tree-like fashion.

// implementation-wise, "things that can emit posts" are termed streamers and are recognized by having a posts() method
// the functions returned by sources in the sources package need to be wrapped in the 'source' struct below in order to give them a posts() method
// Moreover, to enable filters to be added to them, source structs are immediately wrapped in empty stream structs.

// e.g.

// b.Source(s.MySource())

// yields the equivalent of

// b.stream = &stream{
// 	children: []streamer{
// 		&stream{
// 			children: []streamer{
// 				&source{
// 					f: s.MySource(),
// 				},
// 			},
// 		},
// 	}
// }

// s2 := b.Source(s.MySource2())

// adds

// b.stream = &stream{
// 	children: []streamer{
// 		&stream{
// 			children: []streamer{
// 				&source{
// 					f: s.MySource(),
// 				},
// 			},
// 		},
// 		&stream{
// 			children: []streamer{
// 				&source{
// 					f: s.MySource2(),
// 				},
// 			},
// 		},
// 	}
// }

// and assigns that second stream struct to s2
// so that

// s2.Filter(f.S2Filter())

// yields

// b.stream = &stream{
// 	children: []streamer{
// 		&stream{
// 			children: []streamer{
// 				&source{
// 					f: s.MySource(),
// 				},
// 			},
// 		},
// 		&stream{
// 			children: []streamer{
// 				&source{
// 					f: s.MySource2(),
// 				},
// 			},
// 			filters: []func(<-chan content.Post) <-chan content.Post{
// 				f.S2Filter()
// 			}
// 		},
// 	}
// }

// finally, adding another top-level filter

// b.Filter(f.TopFilter())

// would yield

// b.stream = &stream{
// 	children: []streamer{
// 		&stream{
// 			children: []streamer{
// 				&source{
// 					f: s.MySource(),
// 				},
// 			},
// 		},
// 		&stream{
// 			children: []streamer{
// 				&source{
// 					f: s.MySource2(),
// 				},
// 			},
// 			filters: []func(<-chan content.Post) <-chan content.Post{
// 				f.S2Filter()
// 			}
// 		},
// 	},
// 	filters: []func(<-chan content.Post) <-chan content.Post{
// 		f.TopFilter()
// 	}
// }

// as long as the stream struct has a posts() method that aggregates the posts returned by calling posts()
// on each of its children, and then passing those aggregate posts through the filters in its filters array,
// then you can collect the filtered posts of an entire berliner by just calling
// posts() on the top-level stream...

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
	// TODO: cache posts
}

// TODO: abort stuff if it takes too long
// select {
// 	case strChan <- "value":
// 	case <-time.After(5 * time.Second):
// }

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

func clean(f func(*scrape.Client) <-chan content.Post) func(*scrape.Client) <-chan content.Post {
	return func(c *scrape.Client) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range f(c) {
				post.Sanitize()
				out <- post
			}
		}()
		return out
	}
}

func wrapSource(f func(*scrape.Client) <-chan content.Post) *stream {
	return &stream{
		children: []streamer{
			&source{
				f: clean(f),
			},
		},
	}
}
