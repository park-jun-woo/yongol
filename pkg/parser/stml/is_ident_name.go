//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what isIdentName — [A-Za-z_][A-Za-z0-9_]* 식별자 판정 (claims sink 이름 검증)

package stml

// isIdentName reports whether s is a non-empty ASCII identifier
// ([A-Za-z_][A-Za-z0-9_]*) — the shape a claims sink name must have so it
// can be emitted verbatim as a TS object key and a claims map entry.
func isIdentName(s string) bool {
	if s == "" {
		return false
	}
	for i, r := range s {
		isAlpha := r == '_' || (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
		isDigit := r >= '0' && r <= '9'
		if !isAlpha && !(isDigit && i > 0) {
			return false
		}
	}
	return true
}
