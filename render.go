package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
	"github.com/s3ththompson/berliner/browser"
)

var cmdRender = &cobra.Command{
	Use:   "_render",
	Short: "Render articles",
	Long:  "Render an HTML file to stdout from articles and metadata passed in to stdin",
}

func init() {
	cmdRender.Run = Render
}

func Render(cmd *cobra.Command, args []string) {
	posts, err := ReadPosts(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	const html = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Berliner</title>
	<style type="text/css">
		/*! normalize.css v3.0.1 | MIT License | git.io/normalize */html{font-family:sans-serif;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%}body{margin:0}article,aside,details,figcaption,figure,footer,header,hgroup,main,nav,section,summary{display:block}audio,canvas,progress,video{display:inline-block;vertical-align:baseline}audio:not([controls]){display:none;height:0}[hidden],template{display:none}a{background:transparent}a:active,a:hover{outline:0}abbr[title]{border-bottom:1px dotted}b,strong{font-weight:bold}dfn{font-style:italic}h1{font-size:2em;margin:.67em 0}mark{background:#ff0;color:#000}small{font-size:80%}sub,sup{font-size:75%;line-height:0;position:relative;vertical-align:baseline}sup{top:-0.5em}sub{bottom:-0.25em}img{border:0}svg:not(:root){overflow:hidden}figure{margin:1em 40px}hr{-moz-box-sizing:content-box;box-sizing:content-box;height:0}pre{overflow:auto}code,kbd,pre,samp{font-family:monospace,monospace;font-size:1em}button,input,optgroup,select,textarea{color:inherit;font:inherit;margin:0}button{overflow:visible}button,select{text-transform:none}button,html input[type="button"],input[type="reset"],input[type="submit"]{-webkit-appearance:button;cursor:pointer}button[disabled],html input[disabled]{cursor:default}button::-moz-focus-inner,input::-moz-focus-inner{border:0;padding:0}input{line-height:normal}input[type="checkbox"],input[type="radio"]{box-sizing:border-box;padding:0}input[type="number"]::-webkit-inner-spin-button,input[type="number"]::-webkit-outer-spin-button{height:auto}input[type="search"]{-webkit-appearance:textfield;-moz-box-sizing:content-box;-webkit-box-sizing:content-box;box-sizing:content-box}input[type="search"]::-webkit-search-cancel-button,input[type="search"]::-webkit-search-decoration{-webkit-appearance:none}fieldset{border:1px solid #c0c0c0;margin:0 2px;padding:.35em .625em .75em}legend{border:0;padding:0}textarea{overflow:auto}optgroup{font-weight:bold}table{border-collapse:collapse;border-spacing:0}td,th{padding:0}

		body {
		  font-size: 16px;
		  background: #f7f7f7;
		  color: #1d1e19;
		  font-family: "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif;
		  width: 600px;
		  float: right;
		  padding: 20px 30px;
		}

		h1, h2, h3, h4, h5 {
		  font-size: 1em;
		  font-weight: normal;
		  line-height: normal;
		  margin: 0;
		}

		h1 {
		  font-weight: bold;
		}

		h2 {
		  font-style: italic;
		}

		a {
		  text-decoration: none;
		  color: #1d1e19;
		}

		a:hover {
		  text-decoration: underline;
		}

		img {
		  max-width: 100%;
		  max-height: 500px;
		}
	</style>
</head>
<body>
	{{ range .Posts }}
		<article>
			<h1><a href="{{.Permalink}}">{{ .Title }}</a></h1>
			<h2>{{.Authors}}</h2>
			<h2>{{.Source}}</h2>
			<p><img src="{{(index .Images 0).URL}}"></p>
			<div>{{ .Content }}</div>
		</article>
	{{ end }}
</body>
</html>`

	t := template.Must(template.New("html").Parse(html))
	data := struct {
		Posts []*browser.Post
	}{
		posts,
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
