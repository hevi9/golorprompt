package main

// TODO: Change minute coloring based on hour of day section. Section = 24/3

import (
	"github.com/lucasb-eyer/go-colorful"
	//"math"
	"time"
	"fmt"
	"math"
)

func init() {
	SegRegister("time", "Show time",
		func() Segment { return &Time{} })
}

type Time struct{}

func (*Time) Render() []Chunk {
	time1 := time.Now()

	hour := time1.Hour()
	//hourValueScale := math.Min(float64(hour)/23.0001, 1.0)

	hourHue := 200.0 - ((float64(hour % 8)) * (200.0 / 8.0))

	min := time1.Minute()
	minValueScale := math.Min(float64(min)/59.0001, 1.0)

	return []Chunk{
		Chunk{
			text: fmt.Sprintf("%02d", hour),
			fg:   colorful.Hsv(hourHue, config.FgSaturation, config.FgValue),
		},
		Chunk{
			text: ":",
			fg:   config.FgDefault,
		},
		Chunk{
			text: fmt.Sprintf("%02d", min),
			fg:   colorful.Hsv(180.0*minValueScale+180.0, config.FgSaturationLow, config.FgValue),
			//fg:   colorful.Hsv(hourHue, config.FgSaturation, config.FgValue),
		},
	}
}
