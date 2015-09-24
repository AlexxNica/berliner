package sources

import (
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
	"time"
)

func New(name string, entries func(time.Duration) <-chan content.Post) (string, func(*scrape.Client, time.Duration) <-chan content.Post) {
	return name, func(c *scrape.Client, d time.Duration) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			for entry := range entries(d) {
				entry.Via = name
				post, err := c.GetPost(entry.Permalink)
				if err != nil {
					continue
				}
				out <- content.MergePosts(entry, post)
			}
		}()
		return out
	}
}
