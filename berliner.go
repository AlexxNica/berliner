package main

import (
	"github.com/spf13/cobra"
)

func main() {

	Berliner := &cobra.Command{
		Use:   "berliner",
		Short: "Daily digest of online news in a beautiful format",
		Long:  "Daily digest of online news in a beautiful format.",
	}

	cmdFetch := &cobra.Command{
		Use:   "_fetch",
		Short: "Fetch feeds from stdin",
		Long:  "Fetch article permalinks for feeds read from stdin.",
		Run:   Fetch,
	}

	cmdParse := &cobra.Command{
		Use:   "_parse",
		Short: "Parse article permalinks from stdin",
		Long:  "Parse structured articles from permalinks from stdin",
		Run:   Parse,
	}

	Berliner.AddCommand(cmdFetch, cmdParse)

	Berliner.Execute()
}
