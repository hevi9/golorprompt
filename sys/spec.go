package sys

import (
	"unicode/utf8"

	"github.com/lucasb-eyer/go-colorful"
)

const (
	// SegmentEntrySymbolName New segment creation function name
	SegmentEntrySymbolName = "NewWithJson"
)

// Environment Shared interace for running data
type Environment interface {
	// Errors In program execution
	Errors() int
	AddError(error) Environment
}

// Chunk Printed part in prompt
type Chunk struct {
	Text string
	Fg   colorful.Color
	Bg   colorful.Color
}

// Len Unicode characters in chunk
func (c *Chunk) Len() int {
	return utf8.RuneCountInString(c.Text)
}

// Segment Segment interaction interface
type Segment interface {
	Render(env Environment) []Chunk
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

func Register(name string, desc string, newFunc NewWithJSONFunc) {
	segmentRegistry[name] = &segmentInfo{
		name:            name,
		desc:            desc,
		newWithJSONFunc: newFunc,
	}
}
