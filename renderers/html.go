package renderers

import (
	"bufio"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/GeertJohan/go.rice"
	"github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/log"
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
		funcMap := template.FuncMap{
			"fDate": func(t time.Time) string {
				return t.Format("January 2, 2006")
			},
			"insertImage": func(body string, image string) string {
				doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
				if err != nil {
					return body
				}
				if doc.Find("p").Length() == 0 {
					return body + image
				}
				doc.Find("p").First().AfterHtml(image)
				out, _ := doc.Html()
				return out
			},
		}
		t, err := template.New("html").Funcs(funcMap).Parse(html)
		if err != nil {
			log.Println(err)
		}
		data := struct {
			Posts []content.Post
			Style string
			Now   time.Time
		}{
			posts,
			css,
			time.Now(),
		}
		f, err := os.Create(filename)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		err = t.Execute(w, data)
		if err != nil {
			log.Println(err)
		}
		w.Flush()
	}
}
