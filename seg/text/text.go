package text

import (
	"encoding/json"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
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
		func(jsonBuf []byte) (sys.Segment, error) {
			segment := &TextDef{}
			err := json.Unmarshal(jsonBuf, segment)
			if err != nil {
				return nil, err
			}
			log.Debug().
				Str("text", segment.Text).
				Float64("hue", segment.Hue).
				Msg("text args")
			return segment, nil
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
