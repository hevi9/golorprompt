package sys

import (
	"strings"
)

// Widget Interchangeable interface to manage widgets
type Widget interface {
	Render(env Environment, maxLen int)
	Len() int
	Chunks() []Chunk
	Name() string
}

func widgetsLen(widgets []Widget) int {
	length := 0
	for _, w := range widgets {
		length += w.Len()
	}
	return length
}

//////////////////////////////////////////////////////////////////////////////
// segment widget

type segmentWidget struct {
	name    string
	segment Segment
	chunks  []Chunk
}

func (w *segmentWidget) Render(env Environment, maxLen int) {
	w.chunks = w.segment.Render(env)
}

func (w *segmentWidget) Len() int {
	length := 0
	for _, c := range w.chunks {
		length += c.Len()
	}
	return length
}

func (w *segmentWidget) Chunks() []Chunk {
	return w.chunks
}

func (w *segmentWidget) Name() string {
	return w.name
}

//////////////////////////////////////////////////////////////////////////////
// filler

type fillerWidget struct {
	length int
}

func (s *fillerWidget) Allocate(maxLen int) {
	s.length = maxLen + 1 // +1 = the already allocated space
}

func (s *fillerWidget) Len() int {
	return 1
}

func (s *fillerWidget) Chunks() []Chunk {

	return []Chunk{
		Chunk{Text: strings.Repeat(" ", maxInt(s.length, 1))},
	}
}
