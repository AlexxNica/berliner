package sources

import (
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

func GrantlandFeatures() (string, func(scrape.Client, time.Duration) <-chan content.Post) {
	return New("Grantland: Features", RSS("http://grantland.com/features/feed/"))
}
