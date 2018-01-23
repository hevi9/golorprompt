package main

import (
	"os"
	"github.com/lucasb-eyer/go-colorful"
	"log"
	"path"
)

func init() {
	SegRegister("envvar", "Show environment variable if exists",
		func() Segment { return &EnvVar{} })
}

type EnvVar struct {
	Envvar string
	Show   string
}

func (self *EnvVar) Render() []Chunk {
	value, exists := os.LookupEnv(self.Envvar)
	if ! exists {
		return nil
	}
	showFunc, exists := showModifier[self.Show]
	if exists {
		value = showFunc(value)
	} else {
		log.Printf("Error: Show functionality %s does not exists")
	}
	hue := hashToFloat64([]byte(value))
	return []Chunk{
		{text: value, fg: colorful.Hsv(360.0*hue, config.FgSaturation, config.FgValue)},
	}
}

type ShowModifierFunc func(in string) (out string)

var showModifier = map[string]ShowModifierFunc{
	"basename": func(in string) (out string) {
		return path.Base(in)
	},
}
