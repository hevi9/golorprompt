package main

import (
	"log"
	"os"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/sys/unix"
)

func init() {
	SegRegister("cwd", "Current working directory",
		func() Segment { return &Cwd{} })
}

type Cwd struct{}

func (*Cwd) Render() []Chunk {
	var cwd string
	var err error

	if cwd, err = os.Getwd(); err != nil {
		log.Printf("os.Getwd(): %s", err)
		return []Chunk{{text: "NOCWD", fg: config.FgError}}
	}

	parts := strings.Split(cwd, string(os.PathSeparator))
	chunks := make([]Chunk, 0)
	path := ""
	for i := range parts {
		if len(parts[i]) > 0 {
			hue := 330.0*hashToFloat64([]byte(parts[i])) + 15.0
			chunks = append(chunks, Chunk{
				text: parts[i],
				fg:   colorful.Hsv(hue, config.FgSaturationLow, config.FgValue),
			})
		}
		hue := 0.0
		path += parts[i] + string(os.PathSeparator)
		err := unix.Access(path, unix.W_OK)
		if err == nil {
			hue = 90.0
		}
		chunks = append(chunks, Chunk{
			text: string(os.PathSeparator),
			fg:   colorful.Hsv(hue, config.FgSaturation, config.FgValue),
		})

	}
	return chunks
}
