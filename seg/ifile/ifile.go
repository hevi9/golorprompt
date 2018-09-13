package ifile

import (
	"os"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
)

type Ifile struct {
	Filename string
	Sign     string
	Hue      float64
}

func init() {
	sys.Register(
		"ifile",
		"Show sign if file exists",
		func() sys.Segment { return &Ifile{} },
	)
}

func (self *Ifile) Render(env sys.Environment) []sys.Chunk {
	if _, err := os.Stat(self.Filename); os.IsNotExist(err) {
		return nil
	}
	return []sys.Chunk{{
		Text: self.Sign,
		Fg:   colorful.Hsv(self.Hue, sys.Config.FgSaturation, sys.Config.FgValue),
	}}
}
