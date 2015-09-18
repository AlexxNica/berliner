package scrape

import (
	neturl "net/url"
	"strings"
)

func getDomain(url string) string {
	u, err := neturl.Parse(url)
	if err != nil {
		return ""
	}
	parts := strings.Split(u.Host, ".")
	if len(parts) > 2 {
		parts = parts[1:]
	}
	return strings.Join(parts, ".")
}

func isDomain(url string, d string) bool {
	return d == getDomain(url)
}

func must(result string, err error) string {
	if err != nil {
		return ""
	}
	return result
}
