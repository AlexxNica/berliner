package filters

import (
	"github.com/s3ththompson/berliner/content"
	"sort"
)

// Define the sorting interface
type ByPoints []content.Post
func (a ByPoints) Len() int           { return len(a) }
func (a ByPoints) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPoints) Less(i, j int) bool { return a[i].Points < a[j].Points }

// Sort posts in order of points, descending.
//
// Note that because this filter requires knowledge of all posts in the stream
// to sort them, it is a blocking filter that does not output any posts until
// the incoming channel has been closed. This may have performance implications
// depending on where this filter is placed in the chain.
func SortByPoints() (string, func(<-chan content.Post) <-chan content.Post) {
	return "sort by points", func(posts <-chan content.Post) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			allPosts := make([]content.Post, 0)

			for post := range posts {
				allPosts = append(allPosts, post)
			}

			sort.Sort(sort.Reverse(ByPoints(allPosts)))

			for _, post := range allPosts {
				out <- post
			}
		}()
		return out
	}
}
