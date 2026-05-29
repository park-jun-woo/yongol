//ff:func feature=gen-gogin type=util control=sequence
//ff:what parseFuncName — raw func 선언에서 식별자 이름만 추출

package boot

import "strings"

// parseFuncName returns the function identifier from a raw declaration
// such as `func envInt(key string, def int) int { ... }`. Returns "" when
// the prefix is missing or malformed.
func parseFuncName(decl string) string {
	const prefix = "func "
	if !strings.HasPrefix(decl, prefix) {
		return ""
	}
	rest := decl[len(prefix):]
	end := strings.IndexAny(rest, "(<[")
	if end <= 0 {
		return ""
	}
	return strings.TrimSpace(rest[:end])
}
