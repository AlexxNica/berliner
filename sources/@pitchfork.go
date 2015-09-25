package sources

import (
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

func PitchforkBestAlbums() (string, func(scrape.Client, time.Duration) <-chan content.Post) {
	return New("Pitchfork: Best Albums", RSS("http://pitchfork.com/rss/reviews/best/albums/"))
}
