package renderers

import (
	"strconv"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/fatih/color"
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/log"
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
			log.Println(strconv.Itoa(i+1) + ". " + color.WhiteString(post.Title) + " (" + color.GreenString(post.Origin) + ")")
			log.Println("  -> " + color.BlueString(post.Permalink))
			log.Println("  -> " + color.BlueString(strconv.Itoa(post.Points)+" points"))
			log.Println("  -> " + color.BlueString(strconv.Itoa(post.Wordcount())+" words"))
			log.Println("  -> " + color.YellowString(trim(post.Body, 80)+"..."))
		}
	}
}
