package berliner

import (
	"github.com/s3ththompson/berliner/log"
)

type errorWriter struct {
	ch chan log.Entry
}

func newErrorWriter(buffer int) *errorWriter {
	return &errorWriter{
		ch: make(chan log.Entry, buffer),
	}
}

func (w *errorWriter) Write(entry log.Entry) {
	// send entry to channel, or throw away if buffer is full
	select {
		case w.ch <- entry:
		default:
	}
	
}

func (w *errorWriter) entries() <- chan log.Entry {
	return w.ch
}

func (w *errorWriter) close() {
	close(w.ch)
}