package filters

import (
	"github.com/s3ththompson/berliner/content"
)

func contains(posts []content.Post, post content.Post) bool {
	for _, p := range posts {
		if p.Permalink == post.Permalink {
			return true
		}
	}
	return false
}

func Invert(name string, f func(<-chan content.Post) <-chan content.Post) (string, func(<-chan content.Post) <-chan content.Post) {
	return (name + " (inverted)"), func(posts <-chan content.Post) <-chan content.Post {
		filtered := make([]content.Post, 0)
		original := make([]content.Post, 0)
		in := make(chan content.Post)
		go func() {
			defer close(in)
			for post := range posts {
				original = append(original, post)
				in <- post
			}
		}()
		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range f(in) {
				filtered = append(filtered, post)
			}
			for _, post := range original {
				if !contains(filtered, post) {
					out <- post
				}
			}
		}()
		return out
	}
}
