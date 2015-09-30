package sources

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

type hnList []int

type hnItem struct {
	Score int
	Time int64
	Title string
	URL string
}

func getJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func HackerNews50() (string, func(scrape.Client, time.Duration) <-chan content.Post) {
	return New("Hacker News", func(d time.Duration) <-chan content.Post {
		out := make(chan content.Post)
		go func() {
			defer close(out)
			var topItems hnList
			err := getJSON("https://hacker-news.firebaseio.com/v0/topstories.json", &topItems)
			if err != nil {
				return
			}
			topItems = topItems[:50] // TODO: should we arbitrarily limit this to 50?
			var wg sync.WaitGroup
			for _, id := range topItems {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()
					var item hnItem
					url := "https://hacker-news.firebaseio.com/v0/item/"+strconv.Itoa(id)+".json"
					err := getJSON(url, &item)
					if err != nil {
						return
					}
					date := time.Unix(item.Time, 0)
					if d != 0 && time.Since(date) > d {
						return
					}
					out <- content.Post{
						Permalink: item.URL,
						Title: item.Title,
						Date: date,
						Points: item.Score, // TODO: we probably need some conversion here
					}
				}(id)
			}
			wg.Wait()
		}()
		return out
	})
}
