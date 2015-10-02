package filters

import (
	"github.com/s3ththompson/berliner/content"
)

func HasImage() (string, func(<-chan content.Post) <-chan content.Post) {
	return "has image", func(posts <-chan content.Post) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range posts {
				if len(post.Images) > 0 {
					out <- post
				}
			}
		}()
		return out
	}
}
