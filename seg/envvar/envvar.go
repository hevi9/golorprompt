package main

import (
	"log"
	"os"
	"path"

	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	SegRegister("envvar", "Show environment variable if exists",
		func() Segment { return &EnvVar{} })
}

type EnvVar struct {
	Envvar string
	Show   string
	Sign   string
}

// Render ...
func (self *EnvVar) Render() []Chunk {
	value, exists := os.LookupEnv(self.Envvar)
	if !exists {
		return nil
	}
	chunks := make([]Chunk, 0)

	if self.Show != "" {
		showFunc, exists := showModifier[self.Show]
		if exists {
			value = showFunc(self, value)
		} else {
			log.Printf("Error: Show functionality %s does not exists", self.Show)
		}
		hue := hashToFloat64([]byte(value))
		chunks = append(chunks,
			Chunk{text: value, fg: colorful.Hsv(360.0*hue, config.FgSaturation, config.FgValue)})
	}

	if self.Sign != "" {
		hue := hashToFloat64([]byte(self.Sign))
		chunks = append(chunks,
			Chunk{text: self.Sign, fg: colorful.Hsv(360.0*hue, config.FgSaturation, config.FgValue)})
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
