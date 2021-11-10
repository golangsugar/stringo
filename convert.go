package stringo

import (
	"strconv"
	"strings"
)

// AsFloat64 tries to convert a string to float64, and if it can't, just returns zero
func AsFloat64(s string, decimalSeparator, thousandsSeparator rune) float64 {
	if s == "" {
		return 0.0
	}

	s = strings.ReplaceAll(s, string(thousandsSeparator), "")

	s = strings.ReplaceAll(s, string(decimalSeparator), ".")

	f, _ := strconv.ParseFloat(s, 64)

	return f
}
