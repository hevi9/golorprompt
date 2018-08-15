package main

import (
	"os"
	"time"

	"github.com/hevi9/golorprompt/sys"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/alecthomas/kingpin.v2"

	_ "github.com/hevi9/golorprompt/seg/cwd"
)

var (
	cli = kingpin.New("golorprompt", "Generate shell prompt")

	debugFlag = cli.Flag("debug", "Show debug log on console").Bool()

	prompt = cli.Command("prompt", "Generate prompt").Default()

	show = cli.Command("show", "Show information")
)

func main() {
	startTime := time.Now()
	command := kingpin.MustParse(cli.Parse(os.Args[1:]))

	if *debugFlag {
		zerolog.TimeFieldFormat = time.StampMilli
		zerolog.DurationFieldUnit = time.Second
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	app := sys.NewApp()

	switch command {
	case prompt.FullCommand():
		sys.CommandPrompt(app)
	case show.FullCommand():
		sys.CommandShow()
	}

	log.Info().
		Dur("runtime", time.Since(startTime)).
		Int("errors", app.Errors()).
		Msg("done")
}
