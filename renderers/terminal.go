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

func Terminal() (string, func([]content.Post)) {
	return "Terminal", func(posts []content.Post) {
		for i, post := range posts {
			fmt.Println(strconv.Itoa(i+1) + ". " + color.WhiteString(post.Title) + " (" + color.GreenString(post.Origin) + ")")
			fmt.Println("  -> " + color.BlueString(post.Permalink))
			fmt.Println("  -> " + color.BlueString(strconv.Itoa(post.Points) + " points"))
			fmt.Println("  -> " + color.BlueString(strconv.Itoa(post.Wordcount()) + " words"))
			fmt.Println("  -> " + color.YellowString(trim(post.Body, 80)+"..."))
		}
	}
}
