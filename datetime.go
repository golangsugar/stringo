package stringo

import (
	"strings"
	"time"
)

// golangDateFormat translate handy's arbitrary date format to Go's eccentric format
func golangDateTimeFormat(format string) string {
	if format == "" {
		return ""
	}

	newFormat := strings.ToLower(format)

	newFormat = strings.Replace(newFormat, "yyyy", "2006", -1)
	newFormat = strings.Replace(newFormat, "yy", "06", -1)
	newFormat = strings.Replace(newFormat, "mmmm", "January", -1)
	newFormat = strings.Replace(newFormat, "mmm", "Jan", -1)
	newFormat = strings.Replace(newFormat, "mm", "01", -1)
	newFormat = strings.Replace(newFormat, "m", "1", -1)
	newFormat = strings.Replace(newFormat, "dd", "02", -1)
	newFormat = strings.Replace(newFormat, "d", "2", -1)
	newFormat = strings.Replace(newFormat, "hh24", "15", -1)
	newFormat = strings.Replace(newFormat, "hh", "03 PM", -1)
	newFormat = strings.Replace(newFormat, "h", "3 PM", -1)
	newFormat = strings.Replace(newFormat, "nn", "04", -1)
	newFormat = strings.Replace(newFormat, "n", "4", -1)
	newFormat = strings.Replace(newFormat, "ss", "05", -1)
	newFormat = strings.Replace(newFormat, "s", "5", -1)
	newFormat = strings.Replace(newFormat, "ww", "Monday", -1)
	newFormat = strings.Replace(newFormat, "w", "Mon", -1)

	return newFormat
}

// DateTimeAsString formats time.Time variables as strings, considering the format directive
func DateTimeAsString(dt time.Time, format string) string {
	newFormat := golangDateTimeFormat(format)

	return dt.Format(newFormat)
}

// DateReformat gets a date string in a given currentFormat, and transform it according newFormat
func DateReformat(d string, currentLayout, newLayout string) string {
	currentLayout = golangDateTimeFormat(currentLayout)

	if t, err := time.Parse(currentLayout, d); err != nil || t.IsZero() {
		return ""
	} else {
		newLayout = golangDateTimeFormat(newLayout)

		return t.Format(newLayout)
	}
}
