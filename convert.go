package stringo

import (
	"strconv"
	"strings"
)

// AsInt64 tries to convert a string to int64, and if it can't, just returns zero
func AsInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)

	return i
}

// AsInt tries to convert a string to int, and if it can't, just returns zero
func AsInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}

// AsFloat64 tries to convert a string to float64, and if it can't, just returns zero
func AsFloat64(s string, decimalSeparator string) float64 {
	const defaultDecimalSeparator = "."

	thousandSeparator := ","

	if decimalSeparator == "," {
		thousandSeparator = defaultDecimalSeparator
	}

	s = strings.ReplaceAll(s, thousandSeparator, "")

	if decimalSeparator != "." {
		s = strings.Replace(s, decimalSeparator, defaultDecimalSeparator, 1)
	}

	f, _ := strconv.ParseFloat(s, 64)

	return f
}
