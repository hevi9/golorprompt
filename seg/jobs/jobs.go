package main

import (
	"os"
	"github.com/shirou/gopsutil/process"
	"log"
	"github.com/lucasb-eyer/go-colorful"
)

func init() {
	SegRegister("jobs", "Show jobs under shell",
		func() Segment { return &Jobs{} })
}

type Jobs struct{}

func (self Jobs) Render() []Chunk {
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
	chunks := make([]Chunk, 0)
	for _, child := range children {
		status, err := child.Status()
		if child.Pid == thisProcess.Pid {
			continue
		}
		if err != nil {
			log.Printf("Error: child.Status(): %s", err)
			continue
		}
		chunks = append(chunks, Chunk{
			text: status,
			fg:   colorful.Hsv(statusToHue(status), config.FgSaturation, config.FgValue)})
	}
	if len(chunks) > 0 {
		chunks = append(chunks, Chunk{text: sign.jobs, fg: config.FgDefault})
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
