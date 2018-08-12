package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
)

type EnvVar struct {
	Envvar string
	Show   string
	Sign   string
}

func NewWithJson(jsonBuf []byte) sys.Segment {
	segment := &EnvVar{}
	// TODO have error ++ here
	err := json.Unmarshal(jsonBuf, segment)
	if err != nil {
		return nil
	}
	return segment
}

// Render ...
func (self *EnvVar) Render(env sys.Environment) []sys.Chunk {
	value, exists := os.LookupEnv(self.Envvar)
	if !exists {
		return nil
	}
	chunks := make([]sys.Chunk, 0)

	if self.Show != "" {
		showFunc, exists := showModifier[self.Show]
		if exists {
			value = showFunc(self, value)
		} else {
			log.Error().Str("show", self.Show).Msg("Show function does not exists")
			env.AddError(fmt.Errorf("Show function does not exists"))
		}
		hue := sys.HashToFloat64([]byte(value))
		chunks = append(chunks, sys.Chunk{
			Text: value,
			Fg:   colorful.Hsv(360.0*hue, sys.Config.FgSaturation, sys.Config.FgValue),
		})
	}

	if self.Sign != "" {
		hue := sys.HashToFloat64([]byte(self.Sign))
		chunks = append(chunks,
			sys.Chunk{
				Text: self.Sign,
				Fg:   colorful.Hsv(360.0*hue, sys.Config.FgSaturation, sys.Config.FgValue),
			})
	}

	return chunks
}

type showModifierFunc func(self *EnvVar, value string) (out string)

var showModifier = map[string]showModifierFunc{
	"basename": func(self *EnvVar, value string) (out string) {
		return path.Base(value)
	},
	"asis": func(self *EnvVar, value string) (out string) {
		return value
	},
}
