package scrape

import (
	"net/url"
	"strings"
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

func must(result string, err error) string {
	if err != nil {
		return ""
	}
	return result
}
