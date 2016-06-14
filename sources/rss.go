package sources

import (
	"net/url"
	"time"

	"github.com/SlyMarbo/rss"
	"github.com/s3ththompson/berliner/content"
)

func RSS(feed string) func(time.Duration) <-chan content.Post {
	return func(d time.Duration) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			f, err := rss.Fetch(feed)
			if err != nil {
				return
			}
			for _, item := range f.Items {
				if d != 0 && time.Since(item.Date) > d {
					continue
				}
				permalink := item.Link
				// In some cases the permalink may be found in item ID for some reason.
				if permalink == "" {
					// Do some basic validation that it's a real URL before using item ID
					parsed, err := url.Parse(item.ID)
					if err == nil && parsed.Host != "" {
						permalink = item.ID
					}
				}
				out <- content.Post{
					Permalink: permalink,
					Title:     item.Title,
					Date:      item.Date,
				}
			}
		}()
		return out
	}
}
