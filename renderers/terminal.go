package renderers

import (
	"fmt"
	"strconv"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/fatih/color"
	"github.com/s3ththompson/berliner/content"
)

func trim(s string, length int) string {
	l := len(s)
	if l < length {
		return s[0:l]
	} else {
		return s[0:length]
	}
}

func Terminal() func([]content.Post) {
	return func(posts []content.Post) {
		for i, post := range posts {
			fmt.Println(strconv.Itoa(i+1) + ". " + color.WhiteString(post.Title) + " (" + color.GreenString(post.Source) + ")")
			fmt.Println("  -> " + color.BlueString(post.Permalink))
			fmt.Println("  -> " + color.YellowString(trim(post.Body, 80) + "..."))
		}
	}
}
