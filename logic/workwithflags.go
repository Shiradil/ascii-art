package logic

import (
	"fmt"
	"regexp"
)

var Flags []string = []string{
	"--color",
	"--reverse",
	"--align",
}

func GetFlagData(match [][]string) (value, flagType string, err error) {
	if match == nil {
		return "", "", fmt.Errorf("no such flag or value is uppercased")
	}

	return match[0][2], match[0][1], nil
}

func WorkWithFlags(ascii string, match [][]string) (res string, err error) {
	val, flag, err := GetFlagData(match)

	switch flag {
	case Flags[0]:
		res, err = ColorizeAscii(ascii, val)
	case Flags[1]:
		res = ascii
	case Flags[2]:
		res = ascii
	default:
		return "", fmt.Errorf("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	}

	if err != nil {
		return "", err
	}

	return res, nil
}

func IsFlagOrBanner(args []string) (ind int, match [][]string) { // 1 is flag, 2 is banner, 3 is only ascii art, 0 is error
	regexColor := `^(--[a-z]+)=([a-z]+)$`
	comp := regexp.MustCompile(regexColor)
	match = comp.FindAllStringSubmatch(args[1], -1)
	if match != nil && len(args) <= 2 {
		return 0, nil
	} else if match == nil && len(args) <= 2 {
		return 3, nil
	} else if match != nil {
		return 1, match
	} else if args[2] == "standard" || args[2] == "thinkertoy" || args[2] == "shadow" {
		return 2, nil
	} else {
		return 0, nil
	}
}
