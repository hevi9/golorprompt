package main

import (
	"os"
	"github.com/lucasb-eyer/go-colorful"
	"log"
)

func init() {
	SegRegister("hostname", "Show hostname",
		func() Segment { return &Hostname{} })
}

type Hostname struct {
	Threshold int
}

func (self *Hostname) Render() []Chunk {
	if self.Threshold >= 100 {
		return nil
	}
	if _, exists := os.LookupEnv("SSH_CLIENT"); self.Threshold >= 50 && ! exists {
		return nil
	}
	var hostname string
	var err error
	if hostname, err = os.Hostname(); err != nil {
		log.Printf("os.Hostname(): %s", err)
		return nil
	}
	hue := 360.0 * hashToFloat64([]byte(hostname))
	return []Chunk{
		Chunk{
			text: hostname,
			fg:   colorful.Hsv(hue, config.FgSaturationLow, config.FgValue),
		},
	}
}
