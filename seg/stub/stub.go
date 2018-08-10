package main

import (
	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
)

type Stub struct {
}

func NewWithJson(jsonBuf []byte) sys.Segment {
	return &Stub{}
}

func (*Stub) Render(sys.Environment) []sys.Chunk {
	return []sys.Chunk{
		sys.Chunk{
			Text: "01",
			Fg:   colorful.HappyColor(),
		},
		sys.Chunk{
			Text: "23",
			Fg:   colorful.HappyColor(),
		},
		sys.Chunk{
			Text: "45",
			Fg:   colorful.HappyColor(),
		},
		sys.Chunk{
			Text: "6789",
			Fg:   colorful.HappyColor(),
		},
	}
}
