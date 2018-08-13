package main

import (
	"encoding/json"
	"os"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {}

type Ifile struct {
	Filename string
	Sign     string
	Hue      float64
}

func NewWithJson(jsonBuf []byte) sys.Segment {
	segment := &Ifile{}
	// TODO have error ++ here
	err := json.Unmarshal(jsonBuf, segment)
	if err != nil {
		return nil
	}
	return segment
}

func (self *Ifile) Render(env sys.Environment) []sys.Chunk {
	//log.Printf("Ifile.Render: %s %s %f", self.Filename, self.Sign, self.Hue)
	if _, err := os.Stat(self.Filename); os.IsNotExist(err) {
		return nil
	}
	return []sys.Chunk{{
		Text: self.Sign,
		Fg:   colorful.Hsv(self.Hue, sys.Config.FgSaturation, sys.Config.FgValue),
	}}
}
