package ifile

import (
	"encoding/json"
	"os"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
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
		func(jsonBuf []byte) (sys.Segment, error) {
			segment := &Ifile{}
			err := json.Unmarshal(jsonBuf, segment)
			if err != nil {
				return nil, err
			}
			log.Debug().
				Str("filename", segment.Filename).
				Str("sign", segment.Sign).
				Float64("hue", segment.Hue).
				Msg("ifile args")
			return segment, nil
		},
	)
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
