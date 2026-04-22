//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what firstTokenCapture — captures map에서 "token_" 접두사 키 중 알파벳 최소값 반환
package hurl

import (
	"sort"
	"strings"
)

// firstTokenCapture returns the first (alphabetical) capture key that starts
// with "token_". Empty string when no such capture exists. Used as the
// final fallback by resolveTokenVar when neither the role-specific nor the
// single-role "token" capture is available.
func firstTokenCapture(captures map[string]bool) string {
	var tokens []string
	for k := range captures {
		if strings.HasPrefix(k, "token_") {
			tokens = append(tokens, k)
		}
	}
	if len(tokens) == 0 {
		return ""
	}
	sort.Strings(tokens)
	return tokens[0]
}
