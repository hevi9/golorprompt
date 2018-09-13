package load

import (
	"fmt"
	"math"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
)

type Load struct {
	Threshold int
}

func init() {
	sys.Register(
		"load",
		"Alert high cpu load",
		func() sys.Segment {
			return &Load{}
		},
	)
}

func (self *Load) Render(env sys.Environment) []sys.Chunk {

	cpus, err := cpu.Counts(true)
	if err != nil {
		cpus = 1
	}
	stat, err := load.Avg()
	if err != nil {
		log.Error().Err(err).Msg("load.Avg")
		return nil
	}
	if stat.Load1 < (2.0 * float64(cpus) * float64(self.Threshold) / 100.0) {
		return nil
	}
	loadScale := 1.0 - math.Min(stat.Load1/(2.0*float64(cpus)), 1.0)
	return []sys.Chunk{{
		Text: fmt.Sprintf("%.2f%s", stat.Load1, sys.Sign.Load),
		Fg:   colorful.Hsv(90.0*loadScale, sys.Config.FgSaturation, sys.Config.FgValue),
	}}
}
