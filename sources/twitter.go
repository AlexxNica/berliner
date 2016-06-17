package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/content"
)

const twitterJSON = "https://twitter.com/i/profiles/show/%s/timeline?include_available_features=1&include_entities=1"

func Twitter(username string) func(time.Duration) <-chan content.Post {
	return func(d time.Duration) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)

			resp, err := http.Get(fmt.Sprintf(twitterJSON, username))
			if err != nil {
				return
			}
			defer resp.Body.Close()

			// todo: use proper JSON decoder
			var dat map[string]interface{}
			rawJSON, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}
			if err := json.Unmarshal(rawJSON, &dat); err != nil {
				return
			}

			html := dat["items_html"].(string)
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			if err != nil {
				return
			}

			// todo: distinguish between linked tweets and external links
			doc.Find(".tweet-text").Each(func(i int, s *goquery.Selection) {
				url := s.Find("a").AttrOr("data-expanded-url", "")
				if url != "" {
					out <- content.Post{
						Permalink: url,
					}
				}
			})
		}()
		return out
	}
}
