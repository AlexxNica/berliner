// This is a simple test filter that only lets through a set
// number of articles from each source.

package main

import (
	"fmt"
	"os"
	// "github.com/PuerkitoBio/goquery"
	"github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/spf13/cobra"
)

const PerSourceLimit int = 3

func Filter(cmd *cobra.Command, args []string) {
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
		if counter[p.Source] <= PerSourceLimit {
			output = append(output, p)
		}
	}

	WritePosts(output, os.Stdout)
}
