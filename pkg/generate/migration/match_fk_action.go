//ff:func feature=migration type=util control=selection
//ff:what matchFKAction — 대문자 tail 에서 prefix("ON DELETE"/"ON UPDATE") 뒤 동작 토큰 추출
package migration

import "strings"

// matchFKAction returns the recognised action ("CASCADE" / "SET NULL" /
// "RESTRICT" / "NO ACTION") for the given prefix; "" when none matches.
func matchFKAction(upperTail, prefix string) string {
	switch {
	case strings.Contains(upperTail, prefix+" CASCADE"):
		return "CASCADE"
	case strings.Contains(upperTail, prefix+" SET NULL"):
		return "SET NULL"
	case strings.Contains(upperTail, prefix+" RESTRICT"):
		return "RESTRICT"
	case strings.Contains(upperTail, prefix+" NO ACTION"):
		return "NO ACTION"
	}
	return ""
}
