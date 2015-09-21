package filters

import (
	"github.com/s3ththompson/berliner/content"
)

func FixedPoints(points int) func(<-chan content.Post) <-chan content.Post {
	return func(posts <-chan content.Post) <-chan content.Post {
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