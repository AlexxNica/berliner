package berliner

import (
	. "github.com/s3ththompson/berliner/post"
)

type Berliner struct {
	stream stream
}

func New() Berliner {
	return Berliner{}
}

// func (b Berliner) Go() error {
// 	status := b.ctx.start()
// 	update CLI progress with output of status.tick
// 	write CLI messages with output of status.messages
// }

func (b *Berliner) Renderer(f func([]Post)) {
	// b.stream.addRenderer(f)
}

func (b *Berliner) Filter(f func(<-chan Post) <-chan Post) {
	b.stream.addFilter(f)
}

func (b *Berliner) Source(f func() <-chan Post) *stream {
	return b.stream.addSource(f)
}

func (b *Berliner) Posts() <-chan Post {
	return b.stream.posts()
}