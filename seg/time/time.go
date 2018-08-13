package main

import (
	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"

	"fmt"
	"math"
	"time"
)

func main() {}

type Time struct{}

func NewWithJson(jsonBuf []byte) sys.Segment {
	segment := &Time{}
	return segment
}

func (*Time) Render(env sys.Environment) []sys.Chunk {
	time1 := time.Now()

	hour := time1.Hour()
	//hourValueScale := math.Min(float64(hour)/23.0001, 1.0)

	hourHue := 200.0 - ((float64(hour % 8)) * (200.0 / 8.0))

	min := time1.Minute()
	minValueScale := math.Min(float64(min)/59.0001, 1.0)

	return []sys.Chunk{
		sys.Chunk{
			Text: fmt.Sprintf("%02d", hour),
			Fg:   colorful.Hsv(hourHue, sys.Config.FgSaturation, sys.Config.FgValue),
		},
		sys.Chunk{
			Text: ":",
			Fg:   sys.Config.FgDefault,
		},
		sys.Chunk{
			Text: fmt.Sprintf("%02d", min),
			Fg: colorful.Hsv(
				180.0*minValueScale+180.0,
				sys.Config.FgSaturationLow,
				sys.Config.FgValue),
			//fg:   colorful.Hsv(hourHue, config.FgSaturation, config.FgValue),
		},
	}
}
