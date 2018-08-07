package main

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/lucasb-eyer/go-colorful"
	"math"
	"log"
)

func init() {
	SegRegister("disk", "Alert disk usage",
		func() Segment { return &Disk{} })
}

type Disk struct {
	Threshold int
}

func (self *Disk) Render() []Chunk {
	stat, err := disk.Usage(".")
	if err != nil {
		log.Printf("disk.Usage('.'): %s", err)
		return nil
	}
	if stat.UsedPercent < float64(self.Threshold) {
		return nil
	}
	valueScale := 1.0 - math.Min(stat.UsedPercent/100.0, 1.0)
	return []Chunk{{
			text: fmt.Sprintf("%2.f%%%s", stat.UsedPercent, sign.disk),
			fg:   colorful.Hsv(90.0*valueScale, config.FgSaturation, config.FgValue),
		},
	}
}
