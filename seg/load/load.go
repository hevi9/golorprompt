package main

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/cpu"
	"github.com/lucasb-eyer/go-colorful"
	"math"
	"log"
)

func init() {
	SegRegister("load", "Alert high system load",
		func() Segment { return &Load{} })
}

type Load struct {
	Threshold int
}

func (self *Load) Render() []Chunk {

	cpus, err := cpu.Counts(true)
	if err != nil {
		cpus = 1
	}
	stat, err := load.Avg()
	if err != nil {
		log.Printf("load.Avg(): %s", err)
		return nil
	}
	if stat.Load1 < ( 2.0 * float64(cpus) * float64(self.Threshold) / 100.0 ) {
		return nil
	}
	loadScale := 1.0 - math.Min(stat.Load1/(2.0*float64(cpus)), 1.0)
	return []Chunk{{
		text: fmt.Sprintf("%.2f%s", stat.Load1, sign.load),
		fg:   colorful.Hsv(90.0*loadScale, config.FgSaturation, config.FgValue),
	}}
}
