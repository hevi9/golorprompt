package sys

import (
	"log"
	"strings"
)

type Widget interface {
	Allocate(maxLen int)
	Len() int
	Chunks() []Chunk
}

func widgetsLen(widgets []Widget) int {
	length := 0
	for i := range widgets {
		length += widgets[i].Len()
	}
	return length
}

//////////////////////////////////////////////////////////////////////////////

func SegmentWidget(segment Segment) Widget {
	return &segmentWidget{segment: segment}
}

type segmentWidget struct {
	segment Segment
	chunks  []Chunk
}

func (self *segmentWidget) Allocate(maxLen int) {
	chunks := self.segment.Render()
	length := 0
	j := 0
	for i := len(chunks) - 1; i >= 0; i -= 1 {
		length += chunks[i].Len()
		if length > maxLen {
			break
		}
		j = i
	}
	self.chunks = chunks[j:]
	if self.Len() > maxLen {
		log.Printf("segmentWidget: Allocate: maxLen=%d w.Len()=%d chunks=%v\n",
			maxLen, self.Len(), self.chunks)
	}
}

func (self *segmentWidget) Len() int {
	length := 0
	for _, c := range self.chunks {
		length += c.Len()
	}
	return length
}

func (self *segmentWidget) Chunks() []Chunk {
	return self.chunks
}

//////////////////////////////////////////////////////////////////////////////
// space

func Space() Widget {
	return &spaceWidget{len: 1}
}

type spaceWidget struct {
	len int
}

func (self *spaceWidget) Allocate(maxLen int) {
	if maxLen < 1 {
		self.len = 0
	}
}

func (self *spaceWidget) Len() int {
	return self.len
}

func (self *spaceWidget) Chunks() []Chunk {
	if self.len != 0 {
		return []Chunk{{text: " "}}
	} else {
		return nil
	}
}

//////////////////////////////////////////////////////////////////////////////
// filler

func Filler() Widget {
	return &fillerWidget{}
}

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
		Chunk{text: strings.Repeat(" ", MaxInt(s.length, 1))},
	}
}
