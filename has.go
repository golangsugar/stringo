package stringo

import (
	"unicode"
	"unicode/utf8"
)

// HasOnlyNumbers returns true if the input is entirely numeric
func HasOnlyNumbers(sequence string) bool {
	if utf8.RuneCountInString(sequence) == 0 {
		return false
	}

	for _, r := range sequence {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

// HasOnlyDigits returns true if the input is entirely numeric
// HasOnlyDigits is an alias for HasOnlyNumbers()
func HasOnlyDigits(sequence string) bool {
	return HasOnlyNumbers(sequence)
}

// HasOnlyLetters returns true if the input is entirely composed by letters
func HasOnlyLetters(sequence string) bool {
	if utf8.RuneCountInString(sequence) == 0 {
		return false
	}

	for _, r := range sequence {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}
