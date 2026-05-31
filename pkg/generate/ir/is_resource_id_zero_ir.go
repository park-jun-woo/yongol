//ff:func feature=gen-ir type=util control=selection
//ff:what isResourceIDZeroIR -- ResourceID 원시 표현식이 정적 zero 값인지 판정

package ir

import "strings"

// isResourceIDZeroIR returns true when the raw ResourceID expression is a
// static zero value. Mirrors gogin/ssac/isResourceIDZero.
func isResourceIDZeroIR(expr string) bool {
	s := strings.TrimSpace(expr)
	if s == "" {
		return true
	}
	switch strings.ToLower(s) {
	case "0", `""`, "''", "nil", "null":
		return true
	}
	return false
}
