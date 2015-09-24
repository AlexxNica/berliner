package berliner

import (
	"github.com/s3ththompson/berliner/content"
)

func gather(postCh <-chan content.Post) []content.Post {
	posts := []content.Post{}
	for post := range postCh {
		posts = append(posts, post)
	}
	return posts
}

func emit(posts []content.Post) <-chan content.Post {
	out := make(chan content.Post)
	go func() {
		defer close(out)
		for _, post := range posts {
			out <- post
		}
	}()
	return out
}
