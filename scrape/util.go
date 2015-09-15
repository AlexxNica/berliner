package scrape

import (
	neturl "net/url"
	"strings"
)

func domainMatch(url string, d string) bool {
	u, err := neturl.Parse(url)
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

func must(result string, err error) string {
	if err != nil {
		return ""
	}
	return result
}
