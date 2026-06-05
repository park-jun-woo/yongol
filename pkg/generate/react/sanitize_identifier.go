//ff:func feature=gen-react type=util control=sequence topic=string-convert
//ff:what JS 식별자로 무효한 선행 숫자를 Page 접두사로 살균한다 (BUG-107)

package react

// sanitizeComponentName ensures the result is a valid JS/TS identifier.
// If it begins with a digit, prefix "Page" (e.g. "2faSetup" → "Page2faSetup").
// Empty or already-valid identifiers are returned unchanged.
func sanitizeComponentName(s string) string {
	if len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		return "Page" + s
	}
	return s
}
