package renderers

import (
	"bufio"
	"os"
	"path"
	"text/template"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/GeertJohan/go.rice"
	"github.com/s3ththompson/berliner/content"
)

func HTML(filename string, args ...string) (string, func([]content.Post)) {
	htmlfile := "default"
	cssfile := "normalize"
	if len(args) > 0 {
		htmlfile = args[0]
	}
	if len(args) > 1 {
		cssfile = args[1]
	}
	cssBox, err1 := rice.FindBox(path.Join("assets", "css"))
	htmlBox, err2 := rice.FindBox(path.Join("assets", "templates"))
	if err1 != nil || err2 != nil {
		return "HTML", func(posts []content.Post) {}
	}

	css, err1 := cssBox.String(cssfile + ".css")
	html, err2 := htmlBox.String(htmlfile + ".html")
	if err1 != nil || err2 != nil {
		return "HTML", func(posts []content.Post) {}
	}

	return "HTML", func(posts []content.Post) {
		t := template.Must(template.New("html").Parse(html))
		data := struct {
			Posts []content.Post
			Style string
		}{
			posts,
			css,
		}
		f, err := os.Create(filename)
		if err != nil {
			return
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		err = t.Execute(w, data)
		if err != nil {
			return
		}
		w.Flush()
	}
}
