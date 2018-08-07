package main

import "github.com/lucasb-eyer/go-colorful"

func init() {
	SegRegister("stub", "Show stub segment for development purposes",
		func() Segment { return &Stub{} })
}

type Stub struct{}

func (*Stub) Render() []Chunk {
	return []Chunk{
		Chunk{
			text: "01",
			fg:   colorful.HappyColor(),
		},
		Chunk{
			text: "23",
			fg:   colorful.HappyColor(),
		},
		Chunk{
			text: "45",
			fg:   colorful.HappyColor(),
		},
		Chunk{
			text: "6789",
			fg:   colorful.HappyColor(),
		},
	}
}
