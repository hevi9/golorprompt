package sys

import (
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
)

// Environment Shared interace for running data
type Environment interface {
	Errors() int // Errors In program execution
}

// Segment Segment interaction interface
type Segment interface {
	Render(env Environment) []Chunk
}

type Slot struct {
	name    string         `json: segment`
	adjust  int            `json: adjust`
	bg      colorful.Color `json: bg`
	prefix  string         `json: prefix`
	suffix  string         `json: suffix`
	segment Segment
	chunks  []Chunk
}

// NewWithJSONFunc Function signature to call segment creation
// function from plugin. Not (yet) used as type, just a
// specification.
type NewWithJSONFunc func([]byte) (Segment, error)

type segmentInfo struct {
	name            string
	desc            string
	newWithJSONFunc NewWithJSONFunc
}

var segmentRegistry = map[string]*segmentInfo{}

// Register Add segment to registry
func Register(name string, desc string, newFunc NewWithJSONFunc) {
	_, ok := segmentRegistry[name]
	if ok {
		log.Warn().Str("seg", name).Msg("already exists")
	}
	segmentRegistry[name] = &segmentInfo{
		name:            name,
		desc:            desc,
		newWithJSONFunc: newFunc,
	}
}

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
	adjust  int
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
	return length + w.adjust
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
