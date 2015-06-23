// This file defines all the built-in filters, which are cobra subcommands of
// the parent _filter command.
// Having them be subcommands enables different CLI flags per filter type.

package main

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/browser"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
)

func init() {
	cmdClamp.Run = Clamp
	cmdClamp.Flags().IntVarP(&clampLimit, "limit", "l", 3, "Limit for maximum number of articles per source")
	cmdFilter.AddCommand(cmdClamp)

	cmdDedupe.Run = Dedupe
	cmdDedupe.Flags().BoolVarP(&dedupeTest, "test", "t", false, "Allow for test runs by only reading from the logfile, without writing")
	cmdDedupe.Flags().StringVarP(&dedupeLogPath, "file", "f", "", "Specify the file path to persist permalinks")
	cmdFilter.AddCommand(cmdDedupe)
}

// _filter only exists as a cobra command to group its subcommands nicely.
// Calling it directly returns an error.

var cmdFilter = &cobra.Command{
	Use:   "_filter",
	Short: "Filter articles",
	Long:  "Use a subcommand of filter (e.g. _filter clamp) to filter articles",
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
	output := make([]*browser.Post, 0)

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

// Dedupe filter

var dedupeLogPath string
var dedupeTest bool
var cmdDedupe = &cobra.Command{
	Use:   "dedupe",
	Short: "Avoid re-reading articles",
	Long: `Strip out articles that have previously passed through the filter.
					Persists every permalink that passes through, and doesn't let it
					through again.`,
}

func Dedupe(cmd *cobra.Command, args []string) {
	// Check args
	if dedupeLogPath == "" {
		fmt.Fprintln(os.Stderr,
			"Error: Must specify a dedupe logfile path with --file")
		return
	}

	// Open the dedupe log file
	f, err := os.OpenFile(dedupeLogPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Read log file and populate set of permalinks
	lookup := make(map[string]bool) // Use hash table as a set
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lookup[scanner.Text()] = true
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading log file:", err)
		return
	}

	// Read input posts
	input, err := ReadPosts(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Slice of posts to output
	output := make([]*browser.Post, 0)

	// If a post isn't found in the log file, output it, and record it in the log.
	for _, p := range input {
		if !lookup[p.Permalink] {
			output = append(output, p)
			if !dedupeTest {
				_, err = fmt.Fprintln(f, p.Permalink)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
			}
		}
	}

	WritePosts(output, os.Stdout)
}
