package sources

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/SlyMarbo/rss"
	"github.com/s3ththompson/berliner/content"
)

func RSS(feed string) func() <-chan content.Post {
	return func() <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			f, err := rss.Fetch(feed)
			if err != nil {
				return
			}
			for _, item := range f.Items {
				permalink := item.Link
				if permalink == "" { // TODO: remove?
					permalink = item.ID
				}
				out <- content.Post{
					Permalink: permalink,
					Title:     item.Title,
					Date:      item.Date,
					Via:       feed,
				}
			}
		}()
		return out
	}
}
