package helpers

import (
	"regexp"
)

// Using MustCompile will cause the pattern to be compiled at program startup
// and only used in subsequent runs (much faster than Python)
var addressRegex = regexp.MustCompile(`^(1|[13][a-km-zA-HJ-NP-Z0-9]{25,34}|(bc1|[02-9ac-hj-np-z]{1})[a-zA-HJ-NP-Z0-9]{16,100})$`)

// IsValidAddress checks whether the string is a valid Bitcoin address or not
func IsValidAddress(s string) bool {
	return addressRegex.MatchString(s)
}
