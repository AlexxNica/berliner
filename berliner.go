package main

import (
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
)

func main() {
	Berliner := &cobra.Command{
		Use:   "berliner",
		Short: "Daily digest",
		Long:  "Daily digest of online news in a beautiful format.",
	}

	Berliner.AddCommand(cmdFetch, cmdRender, cmdPocket, cmdFilter)

	Berliner.Execute()
}
