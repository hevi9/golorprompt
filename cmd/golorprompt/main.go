package main

import (
	"os"
	"time"

	"github.com/hevi9/golorprompt/sys"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/alecthomas/kingpin.v2"

	// "load" segments
	_ "github.com/hevi9/golorprompt/seg/cwd"
	_ "github.com/hevi9/golorprompt/seg/disk"
	_ "github.com/hevi9/golorprompt/seg/envvar"
	_ "github.com/hevi9/golorprompt/seg/exitcode"
	_ "github.com/hevi9/golorprompt/seg/hostname"
	_ "github.com/hevi9/golorprompt/seg/ifile"
	_ "github.com/hevi9/golorprompt/seg/jobs"
	_ "github.com/hevi9/golorprompt/seg/level"
	_ "github.com/hevi9/golorprompt/seg/load"
	_ "github.com/hevi9/golorprompt/seg/mem"
	_ "github.com/hevi9/golorprompt/seg/start"
	_ "github.com/hevi9/golorprompt/seg/stub"
	_ "github.com/hevi9/golorprompt/seg/time"
	_ "github.com/hevi9/golorprompt/seg/user"
)

var (
	cli = kingpin.New("golorprompt", "Generate shell prompt")

	debugFlag = cli.Flag("debug", "Show debug log on console").Bool()
	config    = cli.Flag("config", "Prompt configuration spec file as json").String()

	prompt = cli.Command("prompt", "Generate prompt").Default()
	show   = cli.Command("show", "Show information")

	// TODO: args eg. RC
)

func main() {
	startTime := time.Now()
	app := sys.NewApp()
	command := kingpin.MustParse(cli.Parse(os.Args[1:]))

	if *debugFlag {
		zerolog.TimeFieldFormat = time.StampMilli
		zerolog.DurationFieldUnit = time.Second
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Hook(app)
		// hook https://github.com/rs/zerolog/blob/master/log_example_test.go
	} else {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	// TODO: Get config json buf
	// 1.cli flag 2. env var 3. default home loc 4. inbuild default

	// Run command
	switch command {
	case prompt.FullCommand():
		sys.CommandPrompt(app, sys.DefaultConfigJSONBuf)
	case show.FullCommand():
		sys.CommandShow()
	}

	log.Info().
		Dur("runtime", time.Since(startTime)).
		Int("errors", app.Errors()).
		Msg("done")
}
