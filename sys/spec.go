package sys

import (
	"unicode/utf8"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	SegmentEntrySymbolName = "NewWithJson"
)

type Environment interface {
}

type Chunk struct {
	Text string
	Fg   colorful.Color
	Bg   colorful.Color
}

func (c *Chunk) Len() int {
	return utf8.RuneCountInString(c.Text)
}

type Segment interface {
	Render(env Environment) []Chunk
}

type NewSegmentFunc func([]byte) Segment
