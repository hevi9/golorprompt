package stub

import (
	"strconv"
	"time"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/rs/zerolog/log"
)

type Stub struct {
	Size int
	Text string
}

func init() {
	sys.Register(
		"stub",
		"Stub segment for development",
		func() sys.Segment { return &Stub{} },
	)
}

func (s *Stub) Render(sys.Environment) []sys.Chunk {
	delay := 100
	log.Debug().Msg("render")
	time.Sleep(delay * time.Millisecond)
	if s.Text != "" {
		return []sys.Chunk{
			sys.Chunk{
				Text: s.Text,
				Fg:   colorful.HappyColor(),
			},
		}
	} else if s.Size > 0 {
		chunks := make([]sys.Chunk, 0)
		for i := 0; i < s.Size; i++ {
			chunks = append(chunks, sys.Chunk{
				Text: strconv.Itoa(i),
				Fg:   colorful.HappyColor(),
			})
		}
		return chunks
	} else {
		return []sys.Chunk{
			sys.Chunk{
				Text: "01",
				Fg:   colorful.HappyColor(),
			},
			sys.Chunk{
				Text: "23",
				Fg:   colorful.HappyColor(),
			},
			sys.Chunk{
				Text: "45",
				Fg:   colorful.HappyColor(),
			},
			sys.Chunk{
				Text: "67",
				Fg:   colorful.HappyColor(),
			},
			sys.Chunk{
				Text: "89",
				Fg:   colorful.HappyColor(),
			},
		}
	}
	return nil
}
