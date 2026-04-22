//ff:func feature=ssac-parse type=util control=iteration dimension=1
//ff:what 문자열이 Go 리터럴 값인지 확인
package ssac

// IsLiteral checks if a string is a Go literal value (not a variable reference).
// Recognizes: numeric (42, -1, 3.14), boolean (true, false), nil.
func IsLiteral(s string) bool {
	if s == "true" || s == "false" || s == "nil" {
		return true
	}
	if len(s) == 0 {
		return false
	}
	start := 0
	if s[0] == '-' {
		start = 1
	}
	if start >= len(s) {
		return false
	}
	dotSeen := false
	for i := start; i < len(s); i++ {
		if s[i] == '.' && !dotSeen {
			dotSeen = true
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
