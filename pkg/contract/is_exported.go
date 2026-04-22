//ff:func feature=contract type=util control=sequence
//ff:what isExported — 식별자 첫 문자가 대문자인지 확인 (Go exported 규칙)

package contract

import "unicode"

// isExported reports whether name starts with an upper-case letter
// per Go's exported-identifier rule. Empty names fail the check.
func isExported(name string) bool {
	if name == "" {
		return false
	}
	r := []rune(name)[0]
	return unicode.IsUpper(r)
}
