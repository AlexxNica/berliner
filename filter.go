// This file defines all the built-in filters, which are cobra subcommands of
// the parent _filter command.
// Having them be subcommands enables different CLI flags per filter type.

package main

import (
	"fmt"
	"os"
	// "github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
)

func init() {
	cmdFilter.Run = Noop

	cmdClamp.Run = Clamp
	cmdClamp.Flags().IntVarP(&clampLimit, "limit", "l", 3, "Limit for maximum number of articles per source")
	cmdFilter.AddCommand(cmdClamp)
}

// _filter only exists as a cobra command to group its subcommands nicely.
// Calling it directly returns an error.

var cmdFilter = &cobra.Command{
	Use:   "_filter",
	Short: "Filter articles",
	Long:  "Use a subcommand of filter (e.g. _filter clamp) to filter articles",
}

func Noop(cmd *cobra.Command, args []string) {
	panic("_filter must be called with a valid subcommand to specify the filter.")
}

// Clamp filter

var clampLimit int
var cmdClamp = &cobra.Command{
	Use:   "clamp",
	Short: "Limit number of articles",
	Long:  "Limit number of articles downloaded per source",
}

func Clamp(cmd *cobra.Command, args []string) {
	input, err := ReadPosts(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Slice of posts to output
	output := make([]*Post, 0)

	// Map to count how many times each source has been seen
	counter := make(map[string]int)

	for _, p := range input {
		counter[p.Source] += 1
		if counter[p.Source] <= clampLimit {
			output = append(output, p)
		}
	}

	WritePosts(output, os.Stdout)
}
