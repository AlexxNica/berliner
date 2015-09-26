package filters

import (
	"bufio"
	"os"

	"github.com/s3ththompson/berliner/content"
	"github.com/s3ththompson/berliner/log"
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

func appendLine(path string, line string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	_, err = w.WriteString(line + "\n")
	if err != nil {
		return err
	}
	return nil
}

// TODO: make this faster avoiding linear time search
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Don't let through two posts with the same permalink.
// Additionally, if a filename is passed, read previous links from that file and
// don't allow through posts which were output from a previous berliner.
func Dedupe(filenames ...string) (string, func(<-chan content.Post) <-chan content.Post) {
	return "Dedupe", func(posts <-chan content.Post) <-chan content.Post {
		var seen []string
		var err error

		var filename string
		if len(filenames) > 0 {
			filename = filenames[0]
			seen, err = readLines(filename)
			if err != nil {
				log.Println(err)
			}
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
		}()
		return out
	}
}

// Lets all posts through while writing permalinks to a file for later use by
// the dedupe filter.
// This filter should be placed last in your filter chain to accurately record
// which posts were included in a final Berliner.
func PersistPosts(filename string) (string, func(<-chan content.Post) <-chan content.Post) {
	return "Persist Posts", func(posts <-chan content.Post) <-chan content.Post {
		var seen []string
		var err error
		seen, err = readLines(filename)
		if err != nil {
			log.Println(err)
		}

		out := make(chan content.Post)
		go func() {
			defer close(out)
			for post := range posts {
				// We check if each link is already in the file just to avoid unnecessarily
				// writing the same link twice.
				if !stringInSlice(post.Permalink, seen) {
					err := appendLine(filename, post.Permalink)
					if err != nil {
						log.Println(err)
					} else {
						seen = append(seen, post.Permalink)
					}
				}
				out <- post
			}
		}()
		return out
	}
}
