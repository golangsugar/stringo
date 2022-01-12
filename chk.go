package stringo

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// HasNumber returns true if input string contains at least one digit/number
func HasNumber(s string) bool {
	for _, s := range s {
		if unicode.IsNumber(s) {
			return true
		}
	}

	return false
}

// HasLetter returns true if input string contains at least one letter
func HasLetter(s string) bool {
	for _, s := range s {
		if unicode.IsLetter(s) {
			return true
		}
	}

	return false
}

// HasSymbol returns true if input string contains at least one symbol
// If rune is not a space, letter nor a number, it's considered a symbol
func HasSymbol(s string) bool {
	for _, s := range s {
		if unicode.IsSymbol(s) || (!unicode.IsLetter(s) && !unicode.IsNumber(s) && !unicode.IsSpace(s)) {
			return true
		}
	}

	return false
}

type ChkRule uint
type ChkResult int

const (
	// ChkAllowEmpty allows empty string ""
	ChkAllowEmpty ChkRule = 1
	// ChkDenySpaces forbids spaces, tabs, new lines and carriage return
	ChkDenySpaces ChkRule = 2
	// ChkDenyNumbers forbids digits/numbers
	ChkDenyNumbers ChkRule = 4
	// ChkDenyLetters forbids letters
	ChkDenyLetters ChkRule = 8
	// ChkDenySymbols forbids symbols. if it's not a number, letter or space, is considered a symbol
	ChkDenySymbols ChkRule = 16
	// ChkDenyMoreThanOneWord forbids more than one word
	ChkDenyMoreThanOneWord ChkRule = 32
	// ChkDenyUpperCase forbids uppercase letters
	ChkDenyUpperCase ChkRule = 64
	// ChkDenyLowercase forbids lowercase letters
	ChkDenyLowercase ChkRule = 128
	// ChkDenyUnicode forbids non-ASCII characters
	ChkDenyUnicode ChkRule = 256
	// ChkRequireNumbers demands at least 1 number within string
	ChkRequireNumbers ChkRule = 512
	// ChkRequireLetters demands at least 1 letter within string
	ChkRequireLetters ChkRule = 1024
	// ChkRequireSymbols demands at least 1 symbol within string. if it's not a number, letter or space, is considered a symbol
	ChkRequireSymbols ChkRule = 2048
	// ChkRequireMoreThanOneWord demands at least 2 words in given string input
	ChkRequireMoreThanOneWord ChkRule = 4096
	// ChkRequireUpperCase demands at least 1 uppercase letter within string
	ChkRequireUpperCase ChkRule = 8192
	// ChkRequireLowercase demands at least 1 lowercase letter within string
	ChkRequireLowercase ChkRule = 16384

	// ChkOk means "alright"
	ChkOk ChkResult = 0
	// ChkEmptyDenied is self explained
	ChkEmptyDenied ChkResult = -1
	// ChkTooShort is self explained
	ChkTooShort ChkResult = -2
	// ChkTooLong is self explained
	ChkTooLong ChkResult = -4
	// ChkSpaceDenied is self explained
	ChkSpaceDenied ChkResult = -5
	// ChkNumbersDenied is self explained
	ChkNumbersDenied ChkResult = -6
	// ChkLettersDenied is self explained
	ChkLettersDenied ChkResult = -7
	// ChkSymbolsDenied is self explained
	ChkSymbolsDenied ChkResult = -8
	// ChkMoreThanOneWordDenied is self explained
	ChkMoreThanOneWordDenied ChkResult = -9
	// ChkUpperCaseDenied is self explained
	ChkUpperCaseDenied ChkResult = -10
	// ChkLowercaseDenied is self explained
	ChkLowercaseDenied ChkResult = -11
	// ChkUnicodeDenied is self explained
	ChkUnicodeDenied ChkResult = -12

	// ChkNumbersNotFound is self explained
	ChkNumbersNotFound ChkResult = -13
	// ChkLettersNotFound is self explained
	ChkLettersNotFound ChkResult = -14
	// ChkSymbolsNotFound is self explained
	ChkSymbolsNotFound ChkResult = -15
	// ChkMoreThanOneWordNotFound is self explained
	ChkMoreThanOneWordNotFound ChkResult = -16
	// ChkUpperCaseNotFound is self explained
	ChkUpperCaseNotFound ChkResult = -17
	// ChkLowercaseNotFound is self explained
	ChkLowercaseNotFound ChkResult = -18
)

var (
	reEmailFinder = regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
)

// CheckStr validates a string according given complexity rules
// CheckStr first evaluates "Deny" rules, and then "Require" rules.
// minLen=0 means there's no minimum length
// maxLen=0 means there's no maximum length
func CheckStr(seq string, minLen, maxLen uint, rules ChkRule) ChkResult {
	strLen := uint(utf8.RuneCountInString(seq))

	if seq == "" {
		if rules&ChkAllowEmpty == ChkAllowEmpty {
			return ChkOk
		}

		return ChkEmptyDenied
	}

	if strLen < minLen {
		return ChkTooShort
	}

	if maxLen > 0 && strLen > maxLen {
		return ChkTooLong
	}

	if rules&ChkDenySpaces == ChkDenySpaces {
		if strings.ContainsAny(seq, "\n\r\t ") {
			return ChkSpaceDenied
		}
	}

	containsNumbers := HasNumber(seq)

	if rules&ChkDenyNumbers == ChkDenyNumbers {
		if containsNumbers {
			return ChkNumbersDenied
		}
	}

	containsLetters := HasLetter(seq)

	if rules&ChkDenyLetters == ChkDenyLetters {
		if containsLetters {
			return ChkLettersDenied
		}
	}

	containsSymbols := HasSymbol(seq)

	if rules&ChkDenySymbols == ChkDenySymbols {
		if containsSymbols {
			return ChkSymbolsDenied
		}
	}

	containsMoreThanOneWord := len(strings.Fields(seq)) > 1

	if rules&ChkDenyMoreThanOneWord == ChkDenyMoreThanOneWord {
		if containsMoreThanOneWord {
			return ChkMoreThanOneWordDenied
		}
	}

	containsUppercase := func(s string) bool {
		if containsLetters {
			for _, s := range s {
				if unicode.IsUpper(s) {
					return true
				}
			}
		}

		return false
	}(seq)

	if rules&ChkDenyUpperCase == ChkDenyUpperCase {
		if containsUppercase {
			return ChkUpperCaseDenied
		}
	}

	containsLowercase := func(s string) bool {
		if containsLetters {
			for _, s := range s {
				if unicode.IsLower(s) {
					return true
				}
			}
		}

		return false
	}(seq)

	if rules&ChkDenyLowercase == ChkDenyLowercase {
		if containsLowercase {
			return ChkLowercaseDenied
		}
	}

	if rules&ChkDenyUnicode == ChkDenyUnicode {
		for _, s := range seq {
			if s > unicode.MaxASCII {
				return ChkUnicodeDenied
			}
		}
	}

	if rules&ChkRequireNumbers == ChkRequireNumbers {
		if !containsNumbers {
			return ChkNumbersNotFound
		}
	}

	if rules&ChkRequireLetters == ChkRequireLetters {
		if !containsLetters {
			return ChkLettersNotFound
		}
	}

	if rules&ChkRequireSymbols == ChkRequireSymbols {
		if !containsSymbols {
			return ChkSymbolsNotFound
		}
	}

	if rules&ChkRequireMoreThanOneWord == ChkRequireMoreThanOneWord {
		if !containsMoreThanOneWord {
			return ChkMoreThanOneWordNotFound
		}
	}

	if rules&ChkRequireUpperCase == ChkRequireUpperCase {
		if !containsUppercase {
			return ChkUpperCaseNotFound
		}
	}

	if rules&ChkRequireLowercase == ChkRequireLowercase {
		if !containsLowercase {
			return ChkLowercaseNotFound
		}
	}

	return ChkOk
}

// StrContainsEmail returns true if given string contains an email address
func StrContainsEmail(seq string) bool {

	return reEmailFinder.MatchString(seq)
}
