//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what classifySingleToken — 단일 Tailwind 클래스를 토큰 카테고리로 분류
package stml_design

import (
	"strings"
)

// classifySingleToken classifies a single Tailwind class into a token category.
func classifySingleToken(part, fullClass, file string, out *pageTokenRefs) {
	// Handle responsive/state prefixes (e.g. "sm:bg-primary", "hover:text-accent")
	if idx := strings.LastIndex(part, ":"); idx >= 0 {
		part = part[idx+1:]
	}

	// Negative prefix (e.g. "-mt-xs")
	stripped := part
	if strings.HasPrefix(stripped, "-") {
		stripped = stripped[1:]
	}

	if matchColorPrefix(stripped, fullClass, file, out) {
		return
	}
	if matchRoundedPrefix(stripped, fullClass, file, out) {
		return
	}
	if matchSpacingPrefix(stripped, fullClass, file, out) {
		return
	}
	matchFontPrefix(stripped, fullClass, file, out)
}
