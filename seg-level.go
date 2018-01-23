package main

import (
	"os"
	"log"
	"strconv"
	"fmt"
)

func init() {
	SegRegister("level", "Alert high system load",
		func() Segment { return &Level{} })
}

type Level struct{}

// SHLVL=0 ? maybe non-interactive shell
// SHLVL=1 - first interactive shell ?

func (self *Level) Render() []Chunk {
	shlvlStr, exists := os.LookupEnv("SHLVL")
	if ! exists {
		log.Printf("Error: no env var SHLVL")
		return nil
	}
	shlvl, err := strconv.Atoi(shlvlStr)
	if err != nil {
		log.Printf("Error: cannot convert SHLVL=%s to int: %s", shlvlStr, err)
		return nil
	}
	if shlvl < 2 {
		return nil
	}
	return []Chunk{{text: fmt.Sprintf("%d%s", shlvl, sign.level), fg: config.FgWarning}}
}
