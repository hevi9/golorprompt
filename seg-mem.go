package main

import (
	"github.com/shirou/gopsutil/mem"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"math"
	"log"
)

func init() {
	SegRegister("mem", "Alert high system memory usage",
		func() Segment { return &Mem{} })
}

type Mem struct {
	Threshold int
}

func (self *Mem) Render() []Chunk {
	stat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("mem.VirtualMemory(): %s", err)
		return nil
	}
	if stat.UsedPercent < float64(self.Threshold) {
		return nil
	}
	valueScale := 1.0 - math.Min(stat.UsedPercent/100.0, 1.0)
	return []Chunk{
		Chunk{
			text: fmt.Sprintf("%2.f%%%s", stat.UsedPercent, sign.memory),
			fg:   colorful.Hsv(90.0*valueScale, config.FgSaturation, config.FgValue),
		},
	}
}
