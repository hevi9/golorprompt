package mem

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/mem"
)

type Mem struct {
	Threshold int
}

func init() {
	sys.Register(
		"mem",
		"Alert high memory usage",
		func(jsonBuf []byte) (sys.Segment, error) {
			segment := &Mem{}
			err := json.Unmarshal(jsonBuf, segment)
			if err != nil {
				return nil, err
			}
			log.Debug().
				Int("threshold", segment.Threshold).
				Msg("mem args")
			return segment, nil
		},
	)
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
