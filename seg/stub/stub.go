package stub

import (
	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
)

type Stub struct{}

func init() {
	sys.Register(
		"stub",
		"Stub segment for development",
		func(jsonBuf []byte) (sys.Segment, error) {
			segment := &Stub{}
			return segment, nil
		},
	)
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
			Text: "67",
			Fg:   colorful.HappyColor(),
		},
		sys.Chunk{
			Text: "89",
			Fg:   colorful.HappyColor(),
		},
	}
}
