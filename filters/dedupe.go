package filters

import (
	"bufio"
	"github.com/s3ththompson/berliner/content"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(file string, lines []string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Dedupe(filenames ...string) func(<-chan content.Post) <-chan content.Post {
	return func(posts <-chan content.Post) <-chan content.Post {
		var seen []string
		persist := false

		var filename string
		if len(filenames) > 0 {
			persist = true
			filename = filenames[0]
			seen, _ = readLines(filename)
		}

		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range posts {
				if !stringInSlice(post.Permalink, seen) {
					out <- post
					seen = append(seen, post.Permalink)
				}
			}
			if persist {
				writeLines(filename, seen)
			}
		}()
		return out
	}
}
