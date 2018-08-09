package sys

import (
	"unicode/utf8"

	"github.com/lucasb-eyer/go-colorful"
)

type Chunk struct {
	text string
	fg   colorful.Color
}

func (c *Chunk) Len() int {
	return utf8.RuneCountInString(c.text)
}

type Segment interface {
	Render() []Chunk
}
