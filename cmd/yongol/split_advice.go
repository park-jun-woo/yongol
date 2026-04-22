//ff:func feature=cli type=util control=iteration dimension=1
//ff:what splitAdvice — extracts inline advice text after "→ Advice:" from a message string
package main

import "strings"

// splitAdvice extracts a trailing suggestion (after " → Advice:" / " — Advice:" /
// "\n↳ Advice:") so the printer can render it on its own line. Returns the main
// message and the advice text (empty if no advice marker found).
func splitAdvice(msg string) (string, string) {
	for _, sep := range []string{"\n↳ Advice:", " → Advice:", " → Advice:", " — Advice:", " — Advice:"} {
		if i := strings.Index(msg, sep); i >= 0 {
			return strings.TrimSpace(msg[:i]), strings.TrimSpace(msg[i+len(sep):])
		}
	}
	return msg, ""
}
