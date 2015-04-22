package main

import (
	"bufio"
	"fmt"
	"os"
)

func readLines() <-chan string {
	out := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		close(out)
	}()
	return out
}
