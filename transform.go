package stringo

import (
	"strings"
	"unicode/utf8"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TransformFlag uint

const (
	// TransformNone No transformations are ordered. Only constraints maximum length
	// TransformNone turns all other flags OFF.
	TransformNone TransformFlag = 1
	// TransformTrim Trim spaces before and after process the input
	// TransformTrim Trims the string, removing leading and trailing spaces
	TransformTrim TransformFlag = 2
	// TransformLowerCase Makes the string lowercase
	// If case transformation flags are combined, the last one remains, considering the following order: TransformTitleCase, TransformLowerCase and TransformUpperCase.
	TransformLowerCase TransformFlag = 4
	// TransformUpperCase Makes the string uppercase
	// If case transformation flags are combined, the last one remains, considering the following order: TransformTitleCase, TransformLowerCase and TransformUpperCase.
	TransformUpperCase TransformFlag = 8
	// TransformOnlyDigits Removes all non-numeric characters
	TransformOnlyDigits TransformFlag = 16
	// TransformOnlyLetters Removes all non-letter characters
	TransformOnlyLetters TransformFlag = 32
	// TransformOnlyLettersAndDigits Leaves only letters and numbers
	TransformOnlyLettersAndDigits TransformFlag = 64
	// TransformHash After process all other flags, applies SHA256 hashing on string for output
	// 	The routine applies handy.Sha256Hash() on given string
	TransformHash TransformFlag = 128
	// TransformTitleCase Makes the string uppercase
	// If case transformation flags are combined, the last one remains, considering the following order: TransformTitleCase, TransformLowerCase and TransformUpperCase.
	TransformTitleCase TransformFlag = 256
	// TransformRemoveDigits Removes all digit characters, without to touch on any other
	// If combined with TransformOnlyLettersAndDigits, TransformOnlyDigits or TransformOnlyLetters, it's ineffective
	TransformRemoveDigits TransformFlag = 512
)

var caser = cases.Title(language.Und)

// Transform handles a string according given flags/parametrization, as follows:
// The transformations are made in arbitrary order, what can result in unexpected output. If the order matters, use TransformSerially instead.
// If maxLen==0, truncation is skipped
// The last operations are, by order, truncation and trimming.
func Transform(s string, maxLen int, transformFlags TransformFlag) string {
	if s == "" {
		return s
	}

	if transformFlags&TransformNone == TransformNone {
		if maxLen > 0 && utf8.RuneCountInString(s) > maxLen {
			s = string([]rune(s)[:maxLen])
		}

		return s
	}

	if (transformFlags & TransformOnlyLettersAndDigits) == TransformOnlyLettersAndDigits {
		s = OnlyLettersAndNumbers(s)
	}

	if (transformFlags & TransformOnlyDigits) == TransformOnlyDigits {
		s = OnlyDigits(s)
	}

	if (transformFlags & TransformOnlyLetters) == TransformOnlyLetters {
		s = OnlyLetters(s)
	}

	if (transformFlags & TransformRemoveDigits) == TransformRemoveDigits {
		s = RemoveDigits(s)
	}

	// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
	if (transformFlags & TransformTrim) == TransformTrim {
		s = strings.TrimSpace(s)
	}

	if (transformFlags & TransformTitleCase) == TransformTitleCase {
		title := caser.String(s)

		s = strings.ToLower(title)
	}

	if (transformFlags & TransformLowerCase) == TransformLowerCase {
		s = strings.ToLower(s)
	}

	if (transformFlags & TransformUpperCase) == TransformUpperCase {
		s = strings.ToUpper(s)
	}

	if (transformFlags & TransformHash) == TransformHash {
		s = Sha256Hash(s)
	}

	if s == "" {
		return s
	}

	if maxLen > 0 && utf8.RuneCountInString(s) > maxLen {
		s = string([]rune(s)[:maxLen])
	}

	// Have to trim before and after, to avoid issues with string truncation and new leading/trailing spaces
	if (transformFlags & TransformTrim) == TransformTrim {
		s = strings.TrimSpace(s)
	}

	return s
}

// TransformSerially reformat given string according parameters, in the order these params were sent
// Example: TransformSerially("uh lalah 123", 4, TransformOnlyDigits,TransformHash,TransformUpperCase)
//          First remove non-digits, then hashes string and after make it all uppercase.
// If maxLen==0, truncation is skipped
// Truncation is the last operation
func TransformSerially(s string, maxLen int, transformFlags ...TransformFlag) string {
	if s == "" {
		return s
	}

	for _, flag := range transformFlags {
		switch flag {
		case TransformOnlyLettersAndDigits:
			s = OnlyLettersAndNumbers(s)
		case TransformOnlyDigits:
			s = OnlyDigits(s)
		case TransformOnlyLetters:
			s = OnlyLetters(s)
		case TransformTrim:
			s = strings.TrimSpace(s)
		case TransformTitleCase:
			s = strings.ToTitle(s)
		case TransformLowerCase:
			s = strings.ToLower(s)
		case TransformUpperCase:
			s = strings.ToUpper(s)
		case TransformHash:
			s = Sha256Hash(s)
		}
	}

	if maxLen > 0 && utf8.RuneCountInString(s) > maxLen {
		s = string([]rune(s)[:maxLen])
	}

	return s
}
