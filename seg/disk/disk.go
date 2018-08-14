package main

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/disk"
)

func main() { /*dummy*/ }

type Disk struct {
	Threshold int
}

func NewWithJson(jsonBuf []byte) sys.Segment {
	segment := &Disk{}
	// TODO have error ++ here
	err := json.Unmarshal(jsonBuf, segment)
	if err != nil {
		return nil
	}
	return segment
}

func (self *Disk) Render(env sys.Environment) []sys.Chunk {
	stat, err := disk.Usage(".")
	if err != nil {
		log.Error().Err(err).Msg("disk.Usage")
		env.AddError(err)
		return nil
	}
	if stat.UsedPercent < float64(self.Threshold) {
		return nil
	}
	valueScale := 1.0 - math.Min(stat.UsedPercent/100.0, 1.0)
	return []sys.Chunk{{
		Text: fmt.Sprintf("%2.f%%%s", stat.UsedPercent, sys.Sign.Disk),
		Fg: colorful.Hsv(
			90.0*valueScale, sys.Config.FgSaturation, sys.Config.FgValue,
		),
	},
	}
}
