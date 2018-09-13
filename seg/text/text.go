package text

import (
	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {}

type TextDef struct {
	Text string
	Hue  float64
}

func init() {
	sys.Register(
		"text",
		"Show given text with hue",
		func() sys.Segment {
			return &TextDef{}
		},
	)
}

func (t *TextDef) Render(env sys.Environment) []sys.Chunk {
	return []sys.Chunk{
		sys.Chunk{
			Text: t.Text,
			Fg:   colorful.Hsv(t.Hue, sys.Config.FgSaturation, sys.Config.FgValue),
		},
	}
}
