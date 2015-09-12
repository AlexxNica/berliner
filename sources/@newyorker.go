package sources

import (
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

func NewYorker() func(*scrape.Client) <-chan content.Post {
	return New("New Yorker", RSS("http://www.newyorker.com/feed/everything"))
}
