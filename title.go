package stringo

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Title tries to return the given word capitalized
func Title(s string) string {
	c := cases.Title(language.Und)

	return c.String(s)
}
