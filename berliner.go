package main

import (
	"github.com/spf13/cobra"
)

func main() {

	Berliner := &cobra.Command{
		Use:   "berliner",
		Short: "Daily digest",
		Long:  "Daily digest of online news in a beautiful format.",
	}

	cmdFetch := &cobra.Command{
		Use:   "_fetch",
		Short: "Fetch feeds",
		Long:  "Fetch article permalinks to stdout for feeds read from stdin.",
		Run:   Fetch,
	}

	cmdParse := &cobra.Command{
		Use:   "_parse",
		Short: "Parse article permalinks",
		Long:  "Parse structured articles to stdout from permalinks from stdin",
		Run:   Parse,
	}

	cmdRender := &cobra.Command{
		Use:   "_render",
		Short: "Render articles",
		Long:  "Render an HTML file to stdout from articles and metadata passed in to stdin",
		Run:   Render,
	}

	Berliner.AddCommand(cmdFetch, cmdParse, cmdRender)

	Berliner.Execute()
}
