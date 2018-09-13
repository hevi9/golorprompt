package disk

import (
	"fmt"
	"math"

	"github.com/hevi9/golorprompt/sys"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/disk"
)

type Disk struct {
	Threshold int
}

func init() {
	sys.Register(
		"disk",
		"Alert for disk capacity",
		func() sys.Segment {
			return &Disk{}
		},
	)
}

func (d *Disk) Render(env sys.Environment) []sys.Chunk {
	stat, err := disk.Usage(".")
	if err != nil {
		log.Error().Err(err).Msg("disk.Usage")
		return nil
	}
	if stat.UsedPercent < float64(d.Threshold) {
		return nil
	}
	valueScale := 1.0 - math.Min(stat.UsedPercent/100.0, 1.0)
	return []sys.Chunk{
		sys.Chunk{
			Text: fmt.Sprintf("%2.f%s", stat.UsedPercent, sys.Sign.Disk),
			Fg:   colorful.Hsv(90.0*valueScale, sys.Config.FgSaturation, sys.Config.FgValue),
		},
	}
}
