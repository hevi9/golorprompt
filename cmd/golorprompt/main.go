package main

import (
	"io/ioutil"
	"os"
	"path"
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
	_ "github.com/hevi9/golorprompt/seg/stub"
	_ "github.com/hevi9/golorprompt/seg/text"
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
	command := kingpin.MustParse(cli.Parse(os.Args[1:]))

	// setup logging
	zerolog.TimeFieldFormat = time.StampMilli
	zerolog.DurationFieldUnit = time.Second
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Hook(sys.Errors)
	// hook https://github.com/rs/zerolog/blob/master/log_example_test.go
	if !*debugFlag {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	// Find config
	// 1.cli flag 2. env var 3. default home loc 4. inbuild default
	jsonBuf := ConfigFile(*config).Resolve().Read()

	// Run command
	switch command {
	case prompt.FullCommand():
		sys.CommandPrompt(jsonBuf)
	case show.FullCommand():
		sys.CommandShow()
	}

	log.Info().
		Dur("runtime", time.Since(startTime)).
		Int("errors", sys.Errors.Count).
		Msg("done")
}

/*

jsonBuf = ConfigFile(*config).Resolve().Read()

*/

type configFile struct {
	filePath string
}

func ConfigFile(config string) *configFile {
	return &configFile{
		filePath: config,
	}
}

func (c *configFile) Resolve() *configFile {
	if c.filePath != "" {
		return c
	}
	value, ok := os.LookupEnv("GOLORPROMPT_CONFIG")
	if ok {
		c.filePath = value
		return c
	}
	value, ok = os.LookupEnv("HOME")
	if !ok {
		log.Error().Str("envvar", "HOME").Msg("Cannot get home path from envvar")
		c.filePath = ""
		return c
	}
	filePath := path.Join(value, ".config", "golorprompt", "prompt.json")
	if _, err := os.Stat(filePath); os.IsExist(err) {
		c.filePath = filePath
		return c
	}
	c.filePath = filePath
	return c
}

func (c *configFile) Read() []byte {
	if c.filePath != "" {
		log.Info().Str("file", c.filePath).Msg("using config")
		buf, err := ioutil.ReadFile(c.filePath)
		if err != nil {
			cwd, err := os.Getwd()
			if err != nil {
				cwd = "ERR"
			}
			log.Error().Err(err).
				Str("filePath", c.filePath).
				Str("cwd", cwd).
				Msg("Cannot read file")
		} else {
			return buf
		}
	}
	log.Debug().Msg("using DEFAULT config")
	return sys.DefaultConfigJSONBuf
}
