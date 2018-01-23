// TODO: Show marker if defined file exists in CWD

package main

import (
	"os"
	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	SegRegister(
		"ifile", "Show sign if given file exists in CWD",
		func() Segment {
			return &Ifile{}
		},
	)
}

type Ifile struct {
	Filename string
	Sign     string
	Hue      float64
}

func (self *Ifile) Render() []Chunk {
	//log.Printf("Ifile.Render: %s %s %f", self.Filename, self.Sign, self.Hue)
	if _, err := os.Stat(self.Filename); os.IsNotExist(err) {
		return nil
	}
	return []Chunk{
		{text: self.Sign, fg: colorful.Hsv(self.Hue, config.FgSaturation, config.FgValue)},
	}
}
