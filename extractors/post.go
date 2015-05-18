package extractors

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"time"
)

type Post struct {
	Title     string
	Permalink string
	Content   string
	Images    []string
	Movies    []string
	Date      time.Time
	Authors   []string
	Tags      []string
	Source    string
	Language  string
}

const PostHeader = "TITLE\tPERMALINK\tCONTENT\tIMAGES\tMOVIES\tDATE\tAUTHORS\tTAGS\tSOURCE\tLANGUAGE"

func sliceToField(s []string) string {
	for i, item := range s {
		if strings.Contains(item, ",") {
			s[i] = strings.Replace(item, ",", "", -1)
		}
	}
	return strings.Join(s, ",")
}

func (p *Post) String() string {
	fields := []string{
		p.Title,
		p.Permalink,
		p.Content,
		sliceToField(p.Images),
		sliceToField(p.Movies),
		p.Date.Format(time.RFC3339),
		sliceToField(p.Authors),
		sliceToField(p.Tags),
		p.Source,
		p.Language,
	}
	return strings.Join(fields, "\t")
}

type PostReader struct {
	r *csv.Reader
}

func NewPostReader(r io.Reader) *PostReader {
	pr := &PostReader{
		r: csv.NewReader(r),
	}
	pr.r.Comma = '\t'
	pr.r.FieldsPerRecord = 10
	pr.r.LazyQuotes = true
	return pr
}

func (pr *PostReader) Read() (*Post, error) {
	record, err := pr.r.Read()
	if err != nil {
		return nil, err
	}
	date, err := time.Parse(time.RFC3339, record[5])
	if err != nil {
		return nil, err
	}
	return &Post{
		Title:     record[0],
		Permalink: record[1],
		Content:   record[2],
		Images:    strings.Split(record[3], ","),
		Movies:    strings.Split(record[4], ","),
		Date:      date,
		Authors:   strings.Split(record[6], ","),
		Tags:      strings.Split(record[7], ","),
		Source:    record[8],
		Language:  record[9],
	}, nil
}

func (pr *PostReader) ReadAll() (posts []*Post, err error) {
	for {
		post, err := pr.Read()
		if err == io.EOF {
			return posts, nil
		}
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
}

func WritePosts(p []*Post, w io.Writer) error {
	_, err := fmt.Fprintln(w, PostHeader) // write csv index
	if err != nil {
		return err
	}
	for post := range p {
		_, err = fmt.Fprintln(w, post)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadPosts(r io.Reader) ([]*Post, error) {
	reader := NewPostReader(r)
	_, err := reader.r.Read() // discard csv index
	if err != nil {
		return nil, err
	}
	return reader.ReadAll()
}
