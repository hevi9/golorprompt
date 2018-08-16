package start

import (
	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {}

type Start struct{}

func init() {
	sys.Register(
		"start",
		"Show a sign",
		func(jsonBuf []byte) (sys.Segment, error) {
			segment := &Start{}
			return segment, nil
		},
	)
}

func (*Start) Render(env sys.Environment) []sys.Chunk {
	return []sys.Chunk{
		sys.Chunk{
			Text: sys.Sign.Start,
			Fg:   colorful.Hsv(45.0, sys.Config.FgSaturation, sys.Config.FgValue),
		},
	}
}
