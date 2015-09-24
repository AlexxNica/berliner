package log

import (
	"time"
	"fmt"

	"github.com/s3ththompson/berliner/content"
)

// logger is private because the only instance is the global package-level one (`std`)
// TODO: add mutex now that writers can be user-provided
type logger struct {
	ch chan Entry
	writer Writer
}

type Entry struct {
	Post content.Post
	Time time.Time
	Message string
}

type Writer interface {
	Write(Entry)
}

var std = new()

func new() *logger {
	return &logger{
		// buffered channel so that logging doesn't block if there's no reader
		ch: make(chan Entry, 1000),
		writer: &StdOutWriter{},
	}
}

// Only errors are exposed because it's an antipattern for filters to log stuff themselves
// All (non-error) logging is handled by the berliner core framework
func Error(args ...interface{}) {
	(&context{}).Error(args...)
}

func Errorf(format string, args ...interface{}) {
	(&context{}).Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	(&context{}).Errorln(args...)
}

func WithPost(post content.Post) *context {
	return &context{
		post: post,
		hasPost: true,
	}
}

func SetWriter(writer Writer) {
	std.writer = writer
}

func ResetWriter() {
	std.writer = &StdOutWriter{}
}

type context struct {
	post content.Post
	hasPost bool // I'm too lazy to check if the content.Post object is actually empty
}

func (ctx *context) log(msg string) {
	entry := Entry{
		Time: time.Now(),
		Message: msg,
	}
	if ctx.hasPost {
		entry.Post = ctx.post
	}
	std.writer.Write(entry)
}

func (ctx *context) Error(args ...interface{}) {
	ctx.log(fmt.Sprint(args...))
}

func (ctx *context) Errorf(format string, args ...interface{}) {
	ctx.Error(fmt.Sprintf(format, args...))
}

func (ctx *context) Errorln(args ...interface{}) {
	ctx.Error(sprintlnn(args...))
}

// TODO: less hacky implementation??
// sprintlnn => Sprint no newline. This is to get the behavior of how
// fmt.Sprintln where spaces are always added between operands, regardless of
// their type. Instead of vendoring the Sprintln implementation to spare a
// string allocation, we do the simplest thing.
func sprintlnn(args ...interface{}) string {
	msg := fmt.Sprintln(args...)
	return msg[:len(msg)-1]
}

type StdOutWriter struct {
}

// TODO: look up regular log behavior
func (w *StdOutWriter) Write(entry Entry) {
	fmt.Println(entry)
}