package sys

import (
	colorful "github.com/lucasb-eyer/go-colorful"
	runewidth "github.com/mattn/go-runewidth"
	"github.com/rs/zerolog/log"
)

// Chunk Printed part in prompt
type Chunk struct {
	Text     string
	Fg       colorful.Color
	Bg       colorful.Color
	Priority int
}

// ChunkH New chunk with hue
func ChunkH(text string, hue float64) Chunk {
	return Chunk{
		Text: text,
		Fg:   colorful.Hsv(hue, Config.FgSaturation, Config.FgValue),
	}
}

// Len Chunk length as terminal display cells. Some unicode
// characters are 2 cells wide.
func (c *Chunk) Len() int {
	log.Debug().
		Str("text", c.Text).
		Int("width", runewidth.StringWidth(c.Text)).
		Msg("runewidth")
	return runewidth.StringWidth(c.Text)
}
