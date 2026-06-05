//ff:func feature=stml-gen type=util control=sequence topic=string-convert
//ff:what JS 식별자로 무효한 선행 숫자를 Page 접두사로 살균한다 (BUG-107)
package stml

// sanitizeIdentifier ensures the result is a valid JS/TS identifier.
// If it begins with a digit, prefix "Page" (e.g. "2faSetup" → "Page2faSetup").
// Local duplicate of react.sanitizeComponentName — stml is a separate package and
// cannot call the parent helper directly. Keep both rules identical so that the
// function declaration (toComponentName) and import/route references (kebabToPascal)
// produce the same identifier for the same kebab input (BUG-107 / Phase031).
func sanitizeIdentifier(s string) string {
	if len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		return "Page" + s
	}
	return s
}
