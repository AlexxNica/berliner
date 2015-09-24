package sources

import (
	"encoding/json"
	"io/ioutil"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

// TODO: change to FromFile and add FromStdin
func FromJSON(filename string) (string, func(*scrape.Client) <-chan content.Post) {
	return "FromJSON", func(c *scrape.Client) <-chan content.Post {
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
				out <- post
			}
		}()
		return out
	}
}
