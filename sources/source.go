package sources

import (
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

func New(name string, entries func() <-chan content.Post) (string, func(*scrape.Client) <-chan content.Post) {
	return name, func(c *scrape.Client) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			for entry := range entries() {
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
