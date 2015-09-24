package sources

import (
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

func NewYorker() (string, func(scrape.Client, time.Duration) <-chan content.Post) {
	return New("New Yorker", RSS("http://www.newyorker.com/feed/everything"))
}
