package filters

import (
	"strings"
	"time"

	"github.com/s3ththompson/berliner/content"
)

func countWords(text string) int {
	return len(strings.Fields(text)) // TODO: just look at words in paragraph tags
}

func ReadFor(duration time.Duration, wpms ...int) func(<-chan content.Post) <-chan content.Post {
	return func(posts <-chan content.Post) <-chan content.Post {
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
				words := countWords(post.Body)
				if words < maxWords { // this should be fuzzy
					out <- post
					maxWords -= words
				} else {
					return
				}
			}
		}()
		return out
	}
}
