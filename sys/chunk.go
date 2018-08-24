package sys

import (
	colorful "github.com/lucasb-eyer/go-colorful"
	runewidth "github.com/mattn/go-runewidth"
)

// Chunk Printed part in prompt
type Chunk struct {
	Text       string
	ColorFg    colorful.Color
	ColorFgUse bool
	Bg         colorful.Color
	BgUse      bool
	Priority   int
	Reset      bool
}

// Chunk builder

func Text(text string) *Chunk {
	return &Chunk{
		Text:       text,
		ColorFgFg:  Config.FgDefault,
		ColorFgUse: true,
	}
}

func (c *Chunk) Fg(colorful.Color) *Chunk {
	return c
}

// ChunkH New chunk with foreground hue
func ChunkH(text string, hue float64) Chunk {
	return Chunk{
		Text:  text,
		Fg:    colorful.Hsv(hue, Config.FgSaturation, Config.FgValue),
		FgUse: true,
	}
}

// ChunkHS New chunk with foreground hue, saturation
func ChunkHS(text string, hue float64, saturation float64) Chunk {
	return Chunk{
		Text:  text,
		Fg:    colorful.Hsv(hue, saturation, Config.FgValue),
		FgUse: true,
	}
}

// ChunkHSV New chunk with foreground hue, saturation, value
func ChunkHSV(text string, hue float64, saturation float64, value float64) Chunk {
	return Chunk{
		Text:  text,
		Fg:    colorful.Hsv(hue, saturation, value),
		FgUse: true,
	}
}

// Len Chunk length as terminal display cells. Some unicode
// characters are 2 cells wide. runewidth package does not
// cover all fonts
func (c *Chunk) Len() int {
	// log.Debug().
	// 	Str("text", c.Text).
	// 	Int("width", runewidth.StringWidth(c.Text)).
	// 	Msg("runewidth")
	return runewidth.StringWidth(c.Text)
}
