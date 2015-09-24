package sources

import (
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

func DisegnoDaily() (string, func(scrape.Client, time.Duration) <-chan content.Post) {
	return New("Disegno Daily", RSS("http://feeds.feedburner.com/disegnofeed"))
}
