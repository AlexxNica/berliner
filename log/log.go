package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	isatty "github.com/s3ththompson/berliner/Godeps/_workspace/src/github.com/mattn/go-isatty"
)

type Logger struct {
	mu  sync.Mutex
	out io.Writer
	terminal bool
	ids map[int]int
}

var std = New(os.Stdout, isatty.IsTerminal(os.Stdout.Fd()))

func Println(args ...interface{}) *entry {
	return std.Println(args...)
}

func New(out io.Writer, isTerminal bool) *Logger {
	return &Logger{
		out: out,
		ids: make(map[int]int),
		terminal: isTerminal,
	}
}

func (l *Logger) isTerminal() bool {
	return l.terminal
}

func (l *Logger) newEntry(message string) *entry {
	id := len(l.ids)
	line := id // while line/id are same today, we may add capability to insert entries between existing lines, which would break this
	l.ids[id] = line
	return &entry{
		l:       l,
		Message: message,
		Time:    time.Now(),
		id:      id,
	}
}

func (l *Logger) Print(args ...interface{}) *entry {
	s := fmt.Sprint(args...)
	return l.output(s)
}

func (l *Logger) Printf(format string, args ...interface{}) *entry {
	s := fmt.Sprintf(format, args...)
	return l.output(s)
}

func (l *Logger) Println(args ...interface{}) *entry {
	s := fmt.Sprintln(args...)
	return l.output(s)
}

func (l *Logger) output(s string) *entry {
	if len(s) == 0 || s[len(s)-1] != '\n' {
		s += "\n"
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	entry := l.newEntry(s)
	fmt.Fprint(l.out, s)
	return entry
}

type entry struct {
	l       *Logger
	Message string
	Time    time.Time
	id      int
}

func (e *entry) stealCursor() {
	diff := len(e.l.ids) - e.l.ids[e.id] // current line - entry line
	// <ESC>[{diff}A = move cursor up diff rows
	fmt.Fprintf(e.l.out, "%c[%dA", 27, diff)
}

func (e *entry) resetCursor() {
	diff := len(e.l.ids) - e.l.ids[e.id] // current line - entry line
	// <ESC>[{diff}B = move cursor down diff rows
	fmt.Fprintf(e.l.out, "%c[%dB", 27, diff)
}

func (e *entry) Update(args ...interface{}) {
	s := fmt.Sprint(args...)
	e.output(s)
}

func (e *entry) Updatef(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	e.output(s)
}

func (e *entry) Updateln(args ...interface{}) {
	s := fmt.Sprintln(args...)
	e.output(s)
}

func (e *entry) output(s string) {
	if len(s) == 0 || s[len(s)-1] != '\n' {
		s += "\n"
	}
	e.l.mu.Lock()
	defer e.l.mu.Unlock()
	if !e.l.terminal {
		fmt.Fprint(e.l.out, s)
		return
	}
	e.stealCursor()
	fmt.Fprint(e.l.out, s)
	e.resetCursor()
}
