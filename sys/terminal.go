package sys

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/crypto/ssh/terminal"
)

// GetWidth Get terminal width
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

// Fg Foreground color
func Fg(color colorful.Color) string {
	r, g, b := color.RGB255()
	return fmt.Sprintf("%s\x1b[38;2;%d;%d;%dm%s",
		Config.NpcStart, r, g, b, Config.NpcEnd)
}

// Bg Background color
func Bg(color colorful.Color) string {
	r, g, b := color.RGB255()
	return fmt.Sprintf("%s\x1b[48;2;%d;%d;%dm%s",
		Config.NpcStart, r, g, b, Config.NpcEnd)
}

// Rz Reset color/style
func Rz() string {
	return fmt.Sprintf("%s\x1b[0m%s", Config.NpcStart, Config.NpcEnd)
}
