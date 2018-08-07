package main

import (
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
)

func GetWidth() int {
	if width, _, err := terminal.GetSize(int(os.Stdin.Fd())); err == nil {
		return width
	} else if value, exists := os.LookupEnv("COLUMNS"); exists {
		if width, err := strconv.Atoi(value); err == nil {
			return width
		} else {
			return 80
		}
	} else {
		return 80
	}
}

func fg(color colorful.Color) string {
	r, g, b := color.RGB255()
	return fmt.Sprintf("%s\x1b[38;2;%d;%d;%dm%s",
		config.NpcStart, r, g, b, config.NpcEnd)
}

func bg(color colorful.Color) string {
	r, g, b := color.RGB255()
	return fmt.Sprintf("%s\x1b[48;2;%d;%d;%dm%s",
		config.NpcStart, r, g, b, config.NpcEnd)
}

func rz() string {
	return fmt.Sprintf("%s\x1b[0m%s", config.NpcStart, config.NpcEnd)
}
