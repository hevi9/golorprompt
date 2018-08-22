package sys

import "strings"

type shell struct {
	npcStart   string              // non printing character start
	npcEnd     string              // non printing character start
	escapeFunc func(string) string // escape function
}

var zshShell = shell{
	npcStart: "%{",
	npcEnd:   "%}",
	escapeFunc: func(in string) (out string) {
		return strings.Replace(in, "%", "%%", -1)
	},
}

// http://tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html
var bashShell = shell{
	npcStart: "\\[",
	npcEnd:   "\\]",
	escapeFunc: func(in string) (out string) {
		return strings.Replace(in, "\\", "\\\\", -1)
	},
}

var noneShell = shell{
	npcStart: "",
	npcEnd:   "",
	escapeFunc: func(in string) (out string) {
		return in
	},
}
