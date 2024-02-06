package logic

import (
	"fmt"
	"strings"
)

const escape = "\x1b"

const (
	NONE = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
)

func ColorizeAscii(ascii, val string) (coloredAscii string, err error) {
	switch strings.ToUpper(val) {
	case "RED":
		coloredAscii = Format(RED, ascii)
	case "GREEN":
		coloredAscii = Format(GREEN, ascii)
	case "YELLOW":
		coloredAscii = Format(YELLOW, ascii)
	case "BLUE":
		coloredAscii = Format(BLUE, ascii)
	case "PURPLE":
		coloredAscii = Format(PURPLE, ascii)
	default:
		return "", fmt.Errorf("Not supported color")
	}

	return coloredAscii, nil
}

func Format(c int, text string) string {
	return color(c) + text + color(NONE)
}

func color(c int) string {
	if c == NONE {
		return fmt.Sprintf("%s[%dm", escape, c)
	}

	return fmt.Sprintf("%s[3%dm", escape, c)
}
