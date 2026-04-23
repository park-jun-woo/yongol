//ff:func feature=migration type=parser control=sequence
//ff:what normalizeNextval — nextval('s'::regclass) 같은 표현의 내부 ::regclass 캐스트 제거
package migration

import "strings"

// normalizeNextval strips ::regclass / sibling casts from a nextval(...)
// expression's argument.
func normalizeNextval(s string) string {
	open := strings.Index(s, "(")
	closeIdx := strings.LastIndex(s, ")")
	if open < 0 || closeIdx <= open {
		return s
	}
	inner := strings.TrimSpace(s[open+1 : closeIdx])
	if i := strings.LastIndex(inner, "::"); i >= 0 {
		tail := strings.TrimSpace(inner[i+2:])
		if looksLikeCastTarget(tail) {
			inner = strings.TrimSpace(inner[:i])
		}
	}
	return "nextval(" + inner + ")"
}
