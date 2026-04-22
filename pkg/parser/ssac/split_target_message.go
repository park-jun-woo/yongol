//ff:func feature=ssac-parse type=util control=sequence
//ff:what splitTargetMessage — splits a "target \"message\"" string into target and message
package ssac

import "strings"

// splitTargetMessage splits a string of the form `target "message"` into its parts.
func splitTargetMessage(s string) (string, string, string) {
	quoteIdx := strings.IndexByte(s, '"')
	if quoteIdx < 0 {
		return strings.TrimSpace(s), "", ""
	}
	target := strings.TrimSpace(s[:quoteIdx])
	msg, remainder := extractQuoted(s[quoteIdx:])
	return target, msg, strings.TrimSpace(remainder)
}
