package browser

import (
	"testing"
)

func AddImageTest(t *testing.T) {
	p := &Post{}
	p.addImage("http://example.com/image.png", "")
	if len(p.Images) != 1 {
		t.Errorf("Expected length of Post Images array to be 1, got %d", len(p.Images))
	}
}

func AddVideoTest(t *testing.T) {
	p := &Post{}
	p.addVideo("http://example.com/movie.mp4", "")
	if len(p.Videos) != 1 {
		t.Errorf("Expected length of Post Videos array to be 1, got %d", len(p.Videos))
	}
}

var validationTests = []struct {
	in  *Post
	out bool
}{
	{&Post{
		Title:     "My Title",
		Permalink: "http://example.com",
		Content:   "Content",
		Images: []Image{Image{
			URL: "http://example.com/image.png",
			Alt: "Alt Text",
		}},
		Source: "My Source",
	}, true},
	{&Post{}, false},
	{&Post{
		Title: "My Title",
	}, false},
	{&Post{
		Title:     "My Title",
		Permalink: "Not a URL",
		Content:   "Content",
		Source:    "My Source",
	}, false},
	{&Post{
		Title:     "My Title",
		Permalink: "http://example.com",
		Content:   "Content",
		Source:    "My Source",
		Images: []Image{Image{
			Alt: "No image permalink",
		}},
	}, false},
}

func TestValidation(t *testing.T) {
	for _, tt := range validationTests {
		if valid := tt.in.validate(); valid != tt.out {
			t.Errorf("for %v, post.validate() => %t, want %t", tt.in, valid, tt.out)
		}
	}
}

var sanitationTests = []struct {
	in  string
	out string
}{
	{"<p>Hello, <b onclick=alert(1337)>World</b>!</p>", "<p>Hello, <b>World</b>!</p>"},
	{"<body>Hello</body>", "Hello"},
	{"<a href=\"http://www.google.com/\"><img src=\"https://ssl.gstatic.com/accounts/ui/logo_2x.png\"/></a>", "<a href=\"http://www.google.com/\" rel=\"nofollow\"><img src=\"https://ssl.gstatic.com/accounts/ui/logo_2x.png\"/></a>"},
	{"Hello<object width=\"640\" height=\"390\"><embed src=\"https://www.youtube.com\"></embed></object>", "Hello"},
}

func SanitizeTest(t *testing.T) {
	for _, tt := range sanitationTests {
		p := &Post{Content: tt.in}
		p.sanitize()
		if p.Content != tt.out {
			t.Errorf("for content: %s, post.sanitize() => %s, want %s", tt.in, p.Content, tt.out)
		}
	}
}
