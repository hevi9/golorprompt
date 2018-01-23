package main

import (
	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	SegRegister("start", "Show prompt start sign",
		func() Segment { return &Start{} })
}

type Start struct{}

func (*Start) Render() []Chunk {
	return []Chunk{
		Chunk{
			text: sign.start,
			fg:   colorful.Hsv(45.0, config.FgSaturation, config.FgValue),
		},
	}
}
