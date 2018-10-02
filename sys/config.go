package sys

import (
	"github.com/lucasb-eyer/go-colorful"
)

// Config golorprompt runtime configuration
var Config = struct {
	Shell           shell
	BgLine          colorful.Color // Line background color
	FgSaturation    float64        // Color saturation of the foreground text
	FgSaturationLow float64        // Low color saturation of the foreground text
	FgValue         float64        // Color value of the foreground text
	BgHue           float64
	BgSaturation    float64
	BgValue         float64
	FgDefault       colorful.Color    // Default color to use
	FgWarning       colorful.Color    // Default color to use for notice or warning info
	FgError         colorful.Color    // Default color for errors
	FgInfo          colorful.Color    // Default color to use informational info
	Args            map[string]string // Cmd line arg map
	Paths           []string          // Paths to find plugins
}{
	Shell:           noneShell,
	BgLine:          colorful.Hsv(45.0, 0.0, 0.35),
	FgSaturation:    0.5,
	FgSaturationLow: 0.2,
	FgValue:         0.9,
	BgHue:           0.0,
	BgSaturation:    0.0,
	BgValue:         0.35,
	FgDefault:       colorful.Hsv(0.0, 0.0, 0.9),
	FgWarning:       colorful.Hsv(45.0, 0.5, 0.9),
	FgError:         colorful.Hsv(0.0, 0.7, 0.9),
	FgInfo:          colorful.Hsv(90.0, 0.5, 0.9),
	Paths:           []string{"./dist/lib/golorprompt"},
}

// DefaultConfigJSONBuf Default prompt configuration
// var DefaultConfigJSONBuf = []byte(`
// [
// 	{ "segment": "cwd", "prefix": "<<" },
// 	{ "segment": "cwd" }
// ]
// `)

// DefaultConfigJSONBuf Default prompt configuration
var DefaultConfigJSONBuf = []byte(`
[
	{ "segment": "text",
	  "text": "" 
	},
	{ "segment": "cwd" },
	{ "segment": "disk",
	  "adjust": 1,
	  "threshold":10
	},
	{ "segment": "envvar", 
	  "envvar":"SHELL", 
	  "show":"asis", 
	  "sign":"@"
	},
	{ "segment": "exitcode" },
	{ "segment": "hostname" },
	{ "segment":  "ifile", 
	  "filename": "Makefile", 
	  "sign":"§",  
	  "hue": 0
	},
	{ "segment": "level" },
	{ "segment": "load", 
	  "threshold": 1
	},
	{ "segment": "mem",
	  "adjust": 1,	  
	  "threshold": 10
	},
	{ "segment": "time" },
	{ "segment": "user" },
	{ "segment": "text",
	  "text": "", 
	  "hue": 45
	}
]
`)

// var DefaultConfigJSONBuf = []byte(`
// [
// 	{ "segment": "text", "args": {"text": ""} },
// 	{ "segment": "space" },
// 	{ "segment": "cwd" },
// 	{ "segment": "space" },
// 	{ "segment": "disk",
// 	  "adjust": 1,
// 	  "args": {"threshold":55}
// 	},
// 	{ "segment": "space" },
// 	{ "segment": "envvar", "args": {"envvar":"HOME", "show":"asis", "sign":"@"} },
// 	{ "segment": "space" },
// 	{ "segment": "exitcode" },
// 	{ "segment": "space" },
// 	{ "segment": "hostname" },
// 	{ "segment": "space" },
// 	{ "segment": "ifile", "args": {"filename":"Makefile", "sign":"§", "hue": 0} },
// 	{ "segment": "space" },
// 	{ "segment": "level" },
// 	{ "segment": "space" },
// 	{ "segment": "load", "args": {"threshold": 5} },
// 	{ "segment": "space" },
// 	{ "segment": "mem",
// 	  "adjust": 1,
// 	  "args": {
// 		  "threshold": 10
// 	  }
// 	},
// 	{ "segment": "space" },
// 	{ "segment": "stub" },
// 	{ "segment": "space" },
// 	{ "segment": "time" },
// 	{ "segment": "space" },

// 	{ "segment": "newline",
// 	  "args": {
// 		  "bg": 145050050
// 	  }
// 	},
// 	{ "segment": "user" },
// 	{ "segment": "text", "args": {"text": "", "hue": 45} },
// 	{ "segment": "space" }
// ]
// `)

// var defaultJsonV1 = `
// {
//   "left": [
//     {"seg": "hostname",
//      "threshold": 60 },
//     {"seg": "cwd"},
//     {"seg": "git"},
//     {"seg": "envvar",
// 	 "envvar": "VIRTUAL_ENV",
//      "show": "basename"},
//     {"seg": "envvar",
// 	 "envvar": "PIPENV_ACTIVE",
//      "sign": "pe"},
//     {"seg": "ifile",
//      "filename": "Makefile",
//      "sign": "M",
//      "hue": 200.0},
//     {"seg": "ifile",
//      "filename": "manage.py",
//      "sign": "m",
//      "hue": 300.0},
//     {"seg": "ifile",
//      "filename": ".env",
//      "sign": "e",
//      "hue": 45.0}
//   ],
//   "right": [
//     {"seg": "time"},
//     {"seg": "load", "threshold": 80},
//     {"seg": "mem", "threshold": 80},
//     {"seg": "disk", "threshold": 80},
//     {"seg": "jobs"},
// 	{"seg": "level"},
//     {"seg": "exitcode"}
//   ],
//   "start": [
//     {"seg": "user"},
//     {"seg": "start"}
//   ]
// }
// `
