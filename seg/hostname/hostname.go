package main

import (
	"encoding/json"
	"os"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
)

type Hostname struct {
	ShowIfEnv string
}

func NewWithJson(jsonBuf []byte) sys.Segment {
	segment := &Hostname{
		ShowIfEnv: "SSH_CLIENT",
	}
	json.Unmarshal(jsonBuf, segment)
	return segment
}

func (self *Hostname) Render(sys.Environment) []sys.Chunk {
	if _, exists := os.LookupEnv(self.ShowIfEnv); !exists {
		return nil
	}
	var hostname string
	var err error
	if hostname, err = os.Hostname(); err != nil {
		log.Error().Err(err).Msg("os.Hostname")
		return nil
	}
	hue := 360.0 * sys.HashToFloat64([]byte(hostname))
	return []sys.Chunk{
		sys.Chunk{
			Text: hostname,
			Fg:   colorful.Hsv(hue, sys.Config.FgSaturationLow, sys.Config.FgValue),
		},
	}
}
