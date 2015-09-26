package filters

import (
	"fmt"
	"github.com/s3ththompson/berliner/content"
)

func Clamp(n int) (string, func(<-chan content.Post) <-chan content.Post) {
	return fmt.Sprintf("clamp to %d posts", n), func(posts <-chan content.Post) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			i := 0
			for post := range posts {
				if n > i {
					out <- post
					i++
				}
			}
		}()
		return out
	}
}
