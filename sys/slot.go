package sys

import (
	colorful "github.com/lucasb-eyer/go-colorful"
	runewidth "github.com/mattn/go-runewidth"
)

// Slot Interchangeable interface to manage widgets
type Slot interface {
	Name() string
	Bg() colorful.Color
	Render()
	Len() int
	Chunks() []Chunk
	Slots() []Slot
}

func slotsLen(slots []Slot) int {
	length := 0
	for _, s := range slots {
		length += s.Len()
	}
	return length
}

//////////////////////////////////////////////////////////////////////////////
// baseSlot

type baseSlot struct {
	Mname   string
	Madjust int
	Mbg     colorful.Color
	Mprefix string
	Msuffix string
}

func (b *baseSlot) Name() string {
	return b.Mname
}

func (b *baseSlot) Bg() colorful.Color {
	return b.Mbg
}

//////////////////////////////////////////////////////////////////////////////
// segmentSlot

type segmentSlot struct {
	baseSlot
	segment Segment
	chunks  []Chunk
	slots   []Slot
}

func (s *segmentSlot) Render() {
	s.chunks = s.segment.Render(nullEnvironment)
}

func (s *segmentSlot) Len() int {
	length := 0
	for _, c := range s.chunks {
		length += c.Len()
	}
	length += runewidth.StringWidth(s.Mprefix)
	length += runewidth.StringWidth(s.Msuffix)
	return length + s.Madjust
}

func (s *segmentSlot) Chunks() []Chunk {
	return s.chunks
}

func (s *segmentSlot) Slots() []Slot {
	return s.slots
}

//////////////////////////////////////////////////////////////////////////////
//

// type fillerWidget struct {
// 	length int
// }

// func (s *fillerWidget) Allocate(maxLen int) {
// 	s.length = maxLen + 1 // +1 = the already allocated space
// }

// func (s *fillerWidget) Len() int {
// 	return 1
// }

// func (s *fillerWidget) Chunks() []Chunk {

// 	return []Chunk{
// 		Chunk{Text: strings.Repeat(" ", maxInt(s.length, 1))},
// 	}
// }

// type segmentWidget struct {
// 	name    string
// 	segment Segment
// 	adjust  int
// 	chunks  []Chunk
// }

// func (w *segmentWidget) Render(env Environment, maxLen int) {
// 	w.chunks = w.segment.Render(env)
// }

// func (w *segmentWidget) Len() int {
// 	length := 0
// 	for _, c := range w.chunks {
// 		length += c.Len()
// 	}
// 	return length + w.adjust
// }

// func (w *segmentWidget) Chunks() []Chunk {
// 	return w.chunks
// }

// func (w *segmentWidget) Name() string {
// 	return w.name
// }

// type fillerWidget struct {
// 	length int
// }

// func (s *fillerWidget) Allocate(maxLen int) {
// 	s.length = maxLen + 1 // +1 = the already allocated space
// }

// func (s *fillerWidget) Len() int {
// 	return 1
// }

// func (s *fillerWidget) Chunks() []Chunk {

// 	return []Chunk{
// 		Chunk{Text: strings.Repeat(" ", maxInt(s.length, 1))},
// 	}
// }
