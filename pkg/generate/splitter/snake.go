//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what snake — CamelCase / PascalCase 식별자를 snake_case 로 변환 (약어 구간 유지)
package splitter

import (
	"strings"
	"unicode"
)

// snake converts a Go identifier to snake_case. Consecutive uppercase
// runs are preserved as acronyms (HTTPServer → http_server, APIKey →
// api_key). A digit following a letter stays attached; a letter
// following a digit inserts an underscore only when the preceding letter
// was lowercase — matching the convention used by the rest of yongol.
func snake(s string) string {
	var b strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if needsSnakeBreak(runes, i) {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
