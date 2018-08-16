package sys

import "strings"

func shellEscapeZsh(in string) (out string) {
	// TODO: Set escape config for bash and fish
	out = strings.Replace(in, "%", "%%", -1)
	return out
}

// TODO: None shell shell

// http://tldp.org/HOWTO/Bash-Prompt-HOWTO/bash-prompt-escape-sequences.html
