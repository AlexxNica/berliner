package filters

import (
	"github.com/s3ththompson/berliner/content"
)

// Assigns a fixed point value to all posts in a stream
func FixedPoints(points int) (string, func(<-chan content.Post) <-chan content.Post) {
	return "fixed points", func(posts <-chan content.Post) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range posts {
				post.Points = points
				out <- post
			}
		}()
		return out
	}
}
