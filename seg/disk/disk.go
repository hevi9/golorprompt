package disk

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/hevi9/golorprompt/sys"
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
		func(jsonBuf []byte) (sys.Segment, error) {
			segment := &Disk{}
			err := json.Unmarshal(jsonBuf, segment)
			if err != nil {
				return nil, err
			}
			log.Debug().Int("threshold", segment.Threshold).Msg("disk args")
			return segment, nil
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
		sys.ChunkH(
			fmt.Sprintf("%2.f%%%s", stat.UsedPercent, sys.Sign.Disk),
			90.0*valueScale,
		),
	}
}
