package logic

import (
	"ascii/validation"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ConvertWordsToAsciiArtWithNewLinesStr(words []string, letters [95][8]string) string {
	resStr := ""
	if validation.OnlyHasNewLines(words) {
		words = words[1:]
	}

	for _, wordToRead := range words {
		if wordToRead == "" {

			resStr += "\n"
			continue
		}
		resStr += ConvertInputToAsciiArtStr(wordToRead, letters)
	}
	return resStr
}

func ConvertInputToAsciiArtStr(str string, letters [95][8]string) string {
	resStr := ""
	for i := 0; i < 8; i++ {
		for _, strLetterIndex := range str {
			if strLetterIndex >= 32 && strLetterIndex <= 126 {
				strLetterIndex -= 32
				for _, line := range letters[strLetterIndex][i] {
					resStr += string(line)
				}
			}
		}
		resStr += "\n"
	}

	return resStr
}

func GetAsciiArtLetters(filePath string) ([95][8]string, error) {
	var empty [95][8]string

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return empty, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var letters [95][8]string
	letterIndex := 0
	for scanner.Scan() {

		for i := 0; i < 8; i++ {
			scanner.Scan()
			line := scanner.Text()

			letters[letterIndex][i] = line
		}
		if err != nil {
			fmt.Println(err)
			return empty, err

		}
		letterIndex++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return empty, err

	}

	return letters, nil
}

func GetAsciiArt(word, pathToStyle string) (string, error) {
	word = strings.ReplaceAll(word, "\n", "\\n")

	if word == "" {
		return "", nil
	}

	if !validation.ConsistsOnlyFromAsciiChars(word) {
		return "", fmt.Errorf("input should consist only from ascii table chars")
	}

	wordsSepByNewLine := strings.Split(word, "\\n")

	letters, err := GetAsciiArtLetters(pathToStyle)
	if err != nil {
		log.Fatal(err)
	}

	resStr := ConvertWordsToAsciiArtWithNewLinesStr(wordsSepByNewLine, letters)

	return resStr, nil
}

func PrintAscii(args []string) {
	ind, match := IsFlagOrBanner(args)
	if len(args) > 2 { // if has a flag or banner
		switch ind {
		case 1:
			resStr, err := GetAsciiArt(args[2], "styles/standard.txt")
			if err != nil {
				fmt.Println(err)
				return
			}

			resFlag, err := WorkWithFlags(resStr, match)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Print(resFlag)
			return
		case 2:
			resStr, err := GetAsciiArt(args[1], "styles/"+args[2]+".txt")
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Print(resStr)
			return
		default:
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\n\nOR\n\nUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	} else { // has only one argument
		if ind == 0 {
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\n\nOR\n\nUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
		resStr, err := GetAsciiArt(args[1], "styles/standard.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(resStr)
	}
}
