package sources

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

// TODO: change to FromFile and add FromStdin
func FromJSON(filename string) (string, func(scrape.Client, time.Duration) <-chan content.Post) {
	return "from JSON", func(c scrape.Client, d time.Duration) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			p, err := ioutil.ReadFile(filename)
			if err != nil {
				return
			}
			posts := make([]content.Post, 10)
			err = json.Unmarshal(p, &posts)
			if err != nil {
				return
			}
			for _, post := range posts {
				if d != 0 && time.Since(post.Date) > d {
					continue
				}
				out <- post
			}
		}()
		return out
	}
}
