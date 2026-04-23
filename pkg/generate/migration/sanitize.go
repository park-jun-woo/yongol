//ff:func feature=migration type=util control=iteration dimension=1
//ff:what sanitize — [a-z0-9_] 외의 문자는 `_` 로 치환 (mnemonic 용)
package migration

import "strings"

// sanitize keeps [a-z0-9_] and replaces others with `_`.
func sanitize(s string) string {
	s = strings.ToLower(s)
	b := strings.Builder{}
	for _, r := range s {
		b.WriteRune(sanitizeRune(r))
	}
	return b.String()
}
