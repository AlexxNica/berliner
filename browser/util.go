package browser

import (
	"bytes"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/golang.org/x/net/html"
	"net/url"
	"strings"
	"unicode"
)

func domainMatch(link string, d string) bool {
	u, err := url.Parse(link)
	if err != nil {
		return false
	}
	parts := strings.Split(u.Host, ".")
	if len(parts) > 2 {
		parts = parts[1:]
	}
	domain := strings.Join(parts, ".")
	return domain == d
}

// TODO: actual trim
func trim(s string) string {
	return s
}

// TODO: actual squeeze
func squeeze(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return ' '
		}
		return r
	}, s)

}

func render(page *html.Node) string {
	var b bytes.Buffer
	err := html.Render(&b, page)
	if err != nil {
		return ""
	}
	return b.String()
}
