package jobs

import (
	"log"
	"os"

	"github.com/hevi9/golorprompt/sys"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/shirou/gopsutil/process"
)

type Jobs struct{}

func init() {
	sys.Register(
		"jobs",
		"Show jobs started from this shell",
		func(jsonBuf []byte) (sys.Segment, error) {
			return &Jobs{}, nil
		},
	)
}

func (self Jobs) Render(env sys.Environment) []sys.Chunk {
	thisProcess, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		log.Printf("Error: process.NewProcess(int32(os.Getpid())): %s", err)
		return nil
	}
	parentProcess, err := thisProcess.Parent()
	if err != nil {
		log.Printf("Error: thisProcess.Parent(): %s", err)
		return nil
	}
	children, err := parentProcess.Children()
	if err != nil {
		log.Printf("Error: parentProcess.Children(): %s", err)
		return nil
	}
	chunks := make([]sys.Chunk, 0)
	for _, child := range children {
		status, err := child.Status()
		if child.Pid == thisProcess.Pid {
			continue
		}
		if err != nil {
			log.Printf("Error: child.Status(): %s", err)
			continue
		}
		chunks = append(chunks, sys.Chunk{
			Text: status,
			Fg:   colorful.Hsv(statusToHue(status), sys.Config.FgSaturation, sys.Config.FgValue)})
	}
	if len(chunks) > 0 {
		chunks = append(chunks, sys.Chunk{Text: sys.Sign.Jobs, Fg: sys.Config.FgDefault})
	}
	return chunks
}

func statusToHue(status string) float64 {
	switch status {
	case "R": // Running
		return 90.0
	case "S": // Sleeping
		return 180.0
	case "T": // sTopped
		return 270.0
	case "I": // Idle
		return 180.0 + 30.0
	case "Z": // Zombie
		return 270.0 + 30.0
	case "W": // Wait
		return 30.0
	default:
		log.Printf("Error: Unknown status: %s", status)
		return 0.0
	}
}
