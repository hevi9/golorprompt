package sys

import (
	colorful "github.com/lucasb-eyer/go-colorful"
	runewidth "github.com/mattn/go-runewidth"
)

// Chunk Printed part in prompt
type Chunk struct {
	Text string
	Fg   colorful.Color
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
