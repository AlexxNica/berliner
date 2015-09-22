// +build !go1.5

package gomail

import "github.com/s3ththompson/berliner/Godeps/_workspace/src/gopkg.in/alexcesaro/quotedprintable.v3"

var newQPWriter = quotedprintable.NewWriter

type mimeEncoder struct {
	quotedprintable.WordEncoder
}

var (
	bEncoding = mimeEncoder{quotedprintable.BEncoding}
	qEncoding = mimeEncoder{quotedprintable.QEncoding}
)
