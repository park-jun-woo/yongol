//ff:func feature=migration type=util control=sequence
//ff:what canonIdent — PostgreSQL 식별자 정규화 (quoted 는 원본 보존, 그 외는 소문자화)
package migration

import "strings"

// canonIdent returns the lowercase form of a PostgreSQL identifier,
// preserving the exact casing for "quoted" identifiers.
func canonIdent(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, ",")
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) && len(s) >= 2 {
		return s[1 : len(s)-1]
	}
	return strings.ToLower(s)
}
