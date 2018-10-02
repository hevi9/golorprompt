package sys

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// Environment Shared interace for running data
type Environment interface {
	Errors() int // Errors In program execution
}

type nullEnvironmentType struct{}

func (*nullEnvironmentType) Errors() int {
	return 0
}

var nullEnvironment = &nullEnvironmentType{}

// Segment Segment interaction interface
type Segment interface {
	Render(env Environment) []Chunk
}

// NewWithJSONFunc Function signature to call segment creation
// function from plugin. Not (yet) used as type, just a
// specification.
type NewSegmentFunc func() Segment

type segmentInfo struct {
	name           string
	desc           string
	newSegmentFunc NewSegmentFunc
}

var segmentRegistry = map[string]*segmentInfo{}

// Register Add segment to registry
func Register(name string, desc string, newFunc NewSegmentFunc) {
	_, ok := segmentRegistry[name]
	if ok {
		log.Warn().Str("seg", name).Msg("already exists")
	}
	segmentRegistry[name] = &segmentInfo{
		name:           name,
		desc:           desc,
		newSegmentFunc: newFunc,
	}
}

func NewSegment(name string) (Segment, error) {
	info, ok := segmentRegistry[name]
	if !ok {
		return nil, fmt.Errorf("name '%s' does not exists", name)
	}
	segment := info.newSegmentFunc()
	log.Info().Str("segment", info.name).Msg("new segment")
	return segment, nil
}
