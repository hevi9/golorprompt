package cwd

import (
	"log"
	"os"
	"strings"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/sys/unix"
)

type cwdS struct{}

func init() {
	sys.Register(
		"cwd",
		"Show current working directory",
		func(jsonBuf []byte) (sys.Segment, error) {
			return &cwdS{}, nil
		},
	)
}

func (*cwdS) Render(env sys.Environment) []sys.Chunk {
	var cwd string
	var err error

	if cwd, err = os.Getwd(); err != nil {
		log.Printf("os.Getwd(): %s", err)
		return []sys.Chunk{{Text: "NOCWD", Fg: sys.Config.FgError}}
	}

	parts := strings.Split(cwd, string(os.PathSeparator))
	chunks := make([]sys.Chunk, 0)
	path := ""
	for i := range parts {
		if len(parts[i]) > 0 {
			hue := 330.0*sys.HashToFloat64([]byte(parts[i])) + 15.0
			chunks = append(chunks, sys.Chunk{
				Text: parts[i],
				Fg:   colorful.Hsv(hue, sys.Config.FgSaturationLow, sys.Config.FgValue),
			})
		}
		hue := 0.0
		path += parts[i] + string(os.PathSeparator)
		err := unix.Access(path, unix.W_OK)
		if err == nil {
			hue = 90.0
		}
		chunks = append(chunks, sys.Chunk{
			Text: string(os.PathSeparator),
			Fg:   colorful.Hsv(hue, sys.Config.FgSaturation, sys.Config.FgValue),
		})

	}
	return chunks
}
