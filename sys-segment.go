package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"unicode/utf8"
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

