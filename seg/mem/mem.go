package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/shirou/gopsutil/mem"
)

type Mem struct {
	Threshold int
}

func NewWithJson(jsonBuf []byte) sys.Segment {
	segment := &Mem{}
	// TODO have error ++ here
	err := json.Unmarshal(jsonBuf, segment)
	if err != nil {
		return nil
	}
	return segment
}

func (self *Mem) Render(env sys.Environment) []sys.Chunk {
	stat, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("mem.VirtualMemory(): %s", err)
		return nil
	}
	if stat.UsedPercent < float64(self.Threshold) {
		return nil
	}
	valueScale := 1.0 - math.Min(stat.UsedPercent/100.0, 1.0)
	return []sys.Chunk{
		sys.Chunk{
			Text: fmt.Sprintf("%2.f%%%s", stat.UsedPercent, sys.Sign.Memory),
			Fg: colorful.Hsv(90.0*valueScale,
				sys.Config.FgSaturation,
				sys.Config.FgValue),
		},
	}
}
