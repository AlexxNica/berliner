package filters

import (
	"fmt"
	"time"

	"github.com/s3ththompson/berliner/content"
)

func ReadFor(duration time.Duration, wpms ...int) (string, func(<-chan content.Post) <-chan content.Post) {
	return fmt.Sprintf("read for %d minutes", int(duration/time.Minute)), func(posts <-chan content.Post) <-chan content.Post {
		wpm := 250
		if len(wpms) > 0 {
			wpm = wpms[0]
		}
		minutes := int(duration / time.Minute)
		maxWords := wpm * minutes

		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range posts {
				words := post.Wordcount()
				if words < maxWords { // this should be fuzzy
					out <- post
					maxWords -= words
				}
			}
		}()
		return out
	}
}
