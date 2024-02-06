package validation

func ConsistsOnlyFromAsciiChars(str string) bool {
	for _, strLetterIndex := range str {
		if !(strLetterIndex >= 32 && strLetterIndex <= 126) {
			return false
		}
	}
	return true
}

func OnlyHasNewLines(words []string) bool {
	for _, wordToRead := range words {
		if wordToRead != "" {
			return false
		}
	}
	return true
}
