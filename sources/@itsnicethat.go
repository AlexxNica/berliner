package sources

import (
	"time"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/scrape"
)

func ItsNiceThat() (string, func(scrape.Client, time.Duration) <-chan content.Post) {
	return New("It's Nice That", RSS("http://feeds2.feedburner.com/itsnicethat/SlXC"))
}
