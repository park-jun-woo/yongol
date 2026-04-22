//ff:func feature=cli type=util control=iteration dimension=1
//ff:what splitAdvice — 메시지에서 "→ 권고:" 이후 권고안을 분리
package main

import "strings"

// splitAdvice extracts a trailing suggestion (after " → 권고:" / " — 권고:" /
// "\n↳ 권고:") so the printer can render it on its own line. Returns the main
// message and the advice text (empty if no advice marker found).
func splitAdvice(msg string) (string, string) {
	for _, sep := range []string{"\n↳ 권고:", " → 권고:", " → 권고안:", " — 권고:", " — 권고안:"} {
		if i := strings.Index(msg, sep); i >= 0 {
			return strings.TrimSpace(msg[:i]), strings.TrimSpace(msg[i+len(sep):])
		}
	}
	return msg, ""
}
