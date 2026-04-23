//ff:func feature=gen-hurl type=util control=selection
//ff:what dummyFieldHint — 필드명 힌트로 dummy 값 결정 (email→{{newUuid}}, password/price/rating/url)

package hurl

import "strings"

// dummyFieldHint returns a dummy value based on the field name.
// Returns nil when no hint matches — caller falls back to type-based default.
//
// BUG-015 / Phase003 — fields whose name contains "email" return a
// {{newUuid}}-templated address so smoke reruns don't collide on the
// DDL unique constraint. Matches both format:"email" (caught earlier
// in dummyString) and plain string fields named email/user_email/
// contact_email that lack the format annotation.
func dummyFieldHint(name string) any {
	lower := strings.ToLower(name)
	switch {
	case strings.Contains(lower, "email"):
		return "smoke-{{newUuid}}@example.com"
	case strings.Contains(lower, "password"):
		return "Password1234!"
	case strings.Contains(lower, "price"), strings.Contains(lower, "amount"):
		return 10000
	case strings.Contains(lower, "rating"):
		return 5
	case strings.Contains(lower, "url"):
		return "https://example.com/test"
	}
	return nil
}
