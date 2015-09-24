package sources

import (
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

func DisegnoDaily() (string, func(*scrape.Client) <-chan content.Post) {
	return New("Disegno Daily", RSS("http://feeds.feedburner.com/disegnofeed"))
}
