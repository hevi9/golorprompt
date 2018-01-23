package main

import "strings"

func shellEscape(in string) (out string) {
	// TODO: Set escape config for bash and fish
	out = strings.Replace(in, "%", "%%", -1)
	return out
}