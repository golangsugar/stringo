package stringo

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type ChkPersonNameResult uint8

const (
	// ChkPersonNameOK means the name was validated
	ChkPersonNameOK ChkPersonNameResult = 0
	// ChkPersonNamePolluted The routine only accepts letters, single quotes and spaces
	ChkPersonNamePolluted ChkPersonNameResult = 1
	// ChkPersonNameTooFewWords The function requires at least 2 words
	ChkPersonNameTooFewWords ChkPersonNameResult = 2
	// ChkPersonNameTooShort the sum of all characters must be >= 6
	ChkPersonNameTooShort ChkPersonNameResult = 3
	// ChkPersonNameTooSimple The name rule requires that at least one word
	ChkPersonNameTooSimple ChkPersonNameResult = 4
)

// ChkPersonName returns true if the name contains at least two words, one >= 3 chars and one >=2 chars.
// I understand that this is a particular criteria, but this is the OpenSourceMagic, where you can change and adapt to your own specs.
func ChkPersonName(name string, acceptEmpty bool) ChkPersonNameResult {
	name = strings.TrimSpace(name)

	// If name is empty, AND it's accepted, return ok. Else, cry!
	if name == "" {
		if !acceptEmpty {
			return ChkPersonNameTooShort
		}

		return ChkPersonNameOK
	}

	// Person names doesn't accept other than letters, spaces and single quotes
	for _, r := range name {
		if !unicode.IsLetter(r) && r != ' ' && r != '\'' && r != '-' {
			return ChkPersonNamePolluted
		}
	}

	// A complete name has to be at least 2 words.
	a := strings.Fields(name)

	if len(a) < 2 {
		return ChkPersonNameTooFewWords
	}

	// At least two words, one with 3 chars and other with 2
	found2 := false
	found3 := false

	for _, s := range a {
		if !found3 && utf8.RuneCountInString(s) >= 3 {
			found3 = true
			continue
		}

		if !found2 && utf8.RuneCountInString(s) >= 2 {
			found2 = true
			continue
		}
	}

	if !found2 || !found3 {
		return ChkPersonNameTooSimple
	}

	return ChkPersonNameOK
}

// NameFirstAndLast returns the first and last words/names from the given input, optionally transformed by transformFlags
// Example: handy.NameFirstAndLast("friedrich wilhelm nietzsche", handy.TransformTitleCase) // returns "Friedrich Nietzsche"
func NameFirstAndLast(name string, transformFlags TransformFlag) string {
	name = strings.Replace(name, "\t", ` `, -1)

	if transformFlags != TransformNone {
		name = Transform(name, utf8.RuneCountInString(name), transformFlags)
	}

	name = strings.TrimSpace(name)

	if name == `` {
		return ``
	}

	words := strings.Split(name, ` `)

	wl := len(words)

	if wl <= 0 {
		return ``
	}

	if wl == 1 {
		return words[0]
	}

	return fmt.Sprintf(`%s %s`, words[0], words[wl-1])
}

// NameFirst returns the first word/name from the given input, optionally transformed by transformFlags
// Example: handy.NameFirst("friedrich wilhelm nietzsche", handy.TransformTitleCase) // returns "Friedrich"
func NameFirst(name string, transformFlags TransformFlag) string {
	name = strings.Replace(name, "\t", ` `, -1)

	if transformFlags != TransformNone {
		name = Transform(name, utf8.RuneCountInString(name), transformFlags)
	}

	name = strings.TrimSpace(name)

	if name == `` {
		return ``
	}

	words := strings.Split(name, ` `)

	if len(words) >= 1 {
		return words[0]
	}

	return ``
}

// Initials returns the first and last words/names from the given input.
func Initials(sequence string) string {
	if len(sequence) < 2 {
		return sequence
	}

	if sequence = strings.TrimSpace(sequence); len(sequence) < 2 {
		return sequence
	}

	var (
		getNext  = true
		initials []rune
	)

	for _, r := range sequence {
		if unicode.IsLetter(r) {
			if getNext {
				initials = append(initials, r)
				getNext = false
			}
		} else {
			getNext = true
		}
	}

	return string(initials)
}
