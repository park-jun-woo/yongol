//ff:func feature=gen-gogin type=util control=sequence
//ff:what isLiteral — Go 리터럴(true/false/nil/숫자/따옴표) 여부 판정

package ssac

import (
	"strings"
	"unicode"
)

// isLiteral returns true for Go literals: true, false, numbers, quoted strings.
func isLiteral(s string) bool {
	if s == "true" || s == "false" || s == "nil" {
		return true
	}
	if strings.HasPrefix(s, `"`) {
		return true
	}
	if len(s) > 0 && (unicode.IsDigit(rune(s[0])) || s[0] == '-') {
		return true
	}
	return false
}
