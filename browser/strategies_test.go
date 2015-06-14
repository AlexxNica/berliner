package browser

import (
	"testing"
)

var bySlugTests = []struct {
	in  string
	out string
	ok  bool
}{
	{"new-york-times", "new-york-times", true},
	{"new-yorker", "new-yorker", true},
	{"nonexistant", "", false},
}

func TestBySlug(t *testing.T) {
	for _, tt := range bySlugTests {
		s, ok := strategies.bySlug(tt.in)
		if ok != tt.ok {
			t.Errorf("strategies.bySlug(%v) => _, %v, want _, %v", tt.in, ok, tt.ok)
		}
		if ok && s.slug() != tt.out {
			t.Errorf("strategies.bySlug(%v) => %v, %v, want %v, %v", tt.in, s.slug(), ok, tt.out, tt.ok)
		}
	}
}

var byLinkTests = []struct {
	in  string
	out string
}{
	{"http://www.nytimes.com/2015/06/10/education/out-of-the-books-in-kindergarten-and-into-the-sandbox.html", "new-york-times"},
	{"http://www.newyorker.com/news/amy-davidson/abby-wambach-chuck-blazer-face-of-american-soccer", "new-yorker"},
	{"http://unrecognized.com", ""},
}

func TestByLink(t *testing.T) {
	for _, tt := range byLinkTests {
		s := strategies.byLink(tt.in)
		if s.slug() != tt.out {
			t.Errorf("strategies.byLink(%v) => %v, want %v", tt.in, s.slug(), tt.out)
		}
	}
}
