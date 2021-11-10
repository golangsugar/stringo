package stringo

import "regexp"

var reEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ValidateEmail returns true if the given input is a valid email address
// Observe that ValidateEmail doesn't trim nor sanitize string before check
// See https://tools.ietf.org/html/rfc2822#section-3.4.1 for details about email address anatomy
func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}

	return reEmail.MatchString(email)
}
