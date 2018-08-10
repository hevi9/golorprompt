package sys

import (
	"github.com/lucasb-eyer/go-colorful"
)

// Config golorprompt runtime configuration
var Config = struct {
	NpcStart        string            // Non-Printing Character sequence start
	NpcEnd          string            // Non-Printing Character sequence end
	BgLine          colorful.Color    // Line background color
	FgSaturation    float64           // Color saturation of the foreground text
	FgSaturationLow float64           // Low color saturation of the foreground text
	FgValue         float64           // Color value of the foreground text
	FgDefault       colorful.Color    // Default color to use
	FgWarning       colorful.Color    // Default color to use for notice or warning info
	FgError         colorful.Color    // Default color for errors
	FgInfo          colorful.Color    // Default color to use informational info
	Args            map[string]string // Cmd line arg map
	Paths           []string          // Paths to find plugins
}{
	NpcStart:        "",
	NpcEnd:          "",
	BgLine:          colorful.Hsv(45.0, 0.0, 0.25),
	FgSaturation:    0.5,
	FgSaturationLow: 0.2,
	FgValue:         0.9,
	FgDefault:       colorful.Hsv(0.0, 0.0, 0.9),
	FgWarning:       colorful.Hsv(45.0, 0.5, 0.9),
	FgError:         colorful.Hsv(0.0, 0.7, 0.9),
	FgInfo:          colorful.Hsv(90.0, 0.5, 0.9),
	Paths:           []string{"./dist/lib/golorprompt"},
}

var defaultJson = `
{
  "left": [
    {"seg": "hostname",
     "threshold": 60 },
    {"seg": "cwd"},
    {"seg": "git"},
    {"seg": "envvar",
	 "envvar": "VIRTUAL_ENV",
     "show": "basename"},
    {"seg": "envvar",
	 "envvar": "PIPENV_ACTIVE",
     "sign": "pe"},
    {"seg": "ifile",
     "filename": "Makefile",
     "sign": "M",
     "hue": 200.0},
    {"seg": "ifile",
     "filename": "manage.py",
     "sign": "m",
     "hue": 300.0},
    {"seg": "ifile",
     "filename": ".env",
     "sign": "e",
     "hue": 45.0}
  ],
  "right": [
    {"seg": "time"},
    {"seg": "load", "threshold": 80},
    {"seg": "mem", "threshold": 80},
    {"seg": "disk", "threshold": 80},
    {"seg": "jobs"},
	{"seg": "level"},
    {"seg": "exitcode"}
  ],
  "start": [
    {"seg": "user"},
    {"seg": "start"}
  ]
}
`
