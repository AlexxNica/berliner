package renderers

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/GeertJohan/go.rice"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/gopkg.in/gomail.v2"
	"github.com/s3ththompson/berliner/content"
	"path"
	"text/template"
	"bytes"
)

type EmailParams struct {
	SmtpServer	string
	SmtpPort		int
	SmtpUsername	string
	SmtpPassword	string
	FromAddress	string
	ToAddress	string
}

func Email(params EmailParams, args ...string) func([]content.Post) {
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
		return func(posts []content.Post) {}
	}

	css, err1 := cssBox.String(cssfile + ".css")
	html, err2 := htmlBox.String(htmlfile + ".html")
	if err1 != nil || err2 != nil {
		return func(posts []content.Post) {}
	}

	return func(posts []content.Post) {
		var doc bytes.Buffer

		t := template.Must(template.New("html").Parse(html))
		data := struct {
			Posts []content.Post
			Style string
		}{
			posts,
			css,
		}
		err := t.Execute(&doc, data)
		if err != nil {
			panic(err)
		}

		m := gomail.NewMessage()
		m.SetHeader("From", params.FromAddress)
		m.SetHeader("To", params.ToAddress)
		m.SetHeader("Subject", "Berliner")
		m.SetBody("text/html", doc.String())

		d := gomail.NewPlainDialer(params.SmtpServer, params.SmtpPort, params.SmtpUsername, params.SmtpPassword)

		err = d.DialAndSend(m)

		if err != nil {
		    panic(err)
		}
	}
}
