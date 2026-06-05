//ff:func feature=stml-parse type=util control=sequence dimension=1
//ff:what 문자가 식별자/숫자/enum-literal 내부 문자인지 판별한다
package stml

import "unicode"

// isGuardIdentRune reports whether r is valid inside an identifier/number/enum.
func isGuardIdentRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-'
}
