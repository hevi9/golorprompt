package level

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hevi9/golorprompt/sys"
)

type Level struct{}

func init() {
	sys.Register(
		"level",
		"Show shell level",
		func() sys.Segment { return &Level{} },
	)
}

// SHLVL=0 ? maybe non-interactive shell
// SHLVL=1 - first interactive shell ?

func (self *Level) Render(sys.Environment) []sys.Chunk {
	shlvlStr, exists := os.LookupEnv("SHLVL")
	if !exists {
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
	return []sys.Chunk{
		{Text: fmt.Sprintf("%d%s", shlvl, sys.Sign.Level), Fg: sys.Config.FgWarning},
	}
}
