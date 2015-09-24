package filters

import (
	"github.com/s3ththompson/berliner/content"
)

func Clamp(n int) (string, func(<-chan content.Post) <-chan content.Post) {
	return "clamp", func(posts <-chan content.Post) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			i := 0
			for post := range posts {
				if n > i {
					out <- post
					i++
				} else {
					return
				}
			}
		}()
		return out
	}
}
