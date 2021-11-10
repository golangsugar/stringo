// Package handy is a toolbelt with utilities and helpers like validators, sanitizers and string formatters.
// There are routines to filter strings, convert between types, validate passwords with custom rules, easily format dates and much more.
package stringo

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

//RuneHasSymbol returns true if the given rune contains a symbol
func RuneHasSymbol(ru rune) bool {
	allowedSymbols := "!\"#$%&'()*+´-./:;<=>?@[\\]^_`{|}~"

	for _, r := range allowedSymbols {
		if ru == r {
			return true
		}
	}

	return false
}

// Sha256Hash simply generates a SHA256 hash from the given string
// In case of error, return ""
func Sha256Hash(s string) string {
	h := sha256.New()

	if _, err := h.Write([]byte(s)); err != nil {
		return ""
	}

	sum := h.Sum(nil)

	return fmt.Sprintf("%x", sum)
}

// Truncate limits the length of a given string, trimming or not, according parameters
func Truncate(s string, maxLen int, trim bool) string {
	if s == "" || maxLen<1{
		return s
	}

	if len(s) > maxLen {
		s = s[0:maxLen]
	}

	if trim {
		s = strings.TrimSpace(s)
	}

	return s
}

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

// TrimLen returns the runes count after trim the spaces
func TrimLen(text string) int {
	if text != "" {
		text = strings.TrimSpace(text)

		if text != "" {
			return utf8.RuneCountInString(text)
		}
	}

	return 0
}

// Reverse returns the given string written backwards, with letters reversed.
func Reverse(s string) string {
	if utf8.RuneCountInString(s) < 2 {
		return s
	}

	r := []rune(s)

	buffer := make([]rune, len(r))

	for i, j := len(r)-1, 0; i >= 0; i-- {
		buffer[j] = r[i]
		j++
	}

	return string(buffer)
}

// ReplaceAll keeps replacing until there's no more occurrences to replace.
func ReplaceAll(original string, replacementPairs ...string) string {
	if original == "" {
		return original
	}

	r := strings.NewReplacer(replacementPairs...)

	for {
		result := r.Replace(original)

		if original != result {
			original = result
		} else {
			break
		}
	}

	return original
}