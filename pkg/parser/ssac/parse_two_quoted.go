//ff:func feature=ssac-parse type=util control=sequence
//ff:what parseTwoQuoted — parses two quoted strings of the form "first" "second"
package ssac

import "strings"

// parseTwoQuoted parses a string of the form `"first" "second"` into two quoted parts.
func parseTwoQuoted(s string) (string, string, string) {
	s = strings.TrimSpace(s)
	first, rest := extractQuoted(s)
	second, remainder := extractQuoted(rest)
	return first, second, strings.TrimSpace(remainder)
}
