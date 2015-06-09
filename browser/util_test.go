package browser

import (
	"testing"
)

var domainTests = []struct {
	link string
	d    string
	out  bool
}{
	{
		"http://www.nytimes.com/2015/06/10/education/out-of-the-books-in-kindergarten-and-into-the-sandbox.html",
		"nytimes.com",
		true,
	},
	{
		"http://rss.nytimes.com/c/34625/f/640308/s/470cf54d/sc/28/l/0L0Snytimes0N0C20A150C0A60C0A90Carts0Cdance0Creview0Ecity0Eballet0Edebuts0Ein0Eseason0Efinale0Bhtml0Dpartner0Frss0Gemc0Frss/story01.htm",
		"nytimes.com",
		true,
	},
	{
		"asdf",
		"nytimes.com",
		false,
	},
}

func TestDomainMatch(t *testing.T) {
	for _, tt := range domainTests {
		if m := domainMatch(tt.link, tt.d); m != tt.out {
			t.Errorf("domainMatch(%v, %v) => %v, want %v", tt.link, tt.d, m, tt.out)
		}
	}
}
