//ff:func feature=gen-hurl type=util control=selection
//ff:what dummyString — string format에 따른 dummy 문자열 값 생성 (email은 {{newUuid}} unique)

package hurl

// dummyString returns a dummy string value based on the format.
//
// BUG-015 / Phase003 — fields with format:"email" embed the hurl 4+
// template variable {{newUuid}} so each smoke run produces a fresh,
// unique address. This lets Register succeed repeatedly on the same
// DB (empty DB → 201, second run → 201 again with a different uuid)
// instead of the previous test@example.com which collided on the
// unique-email DDL constraint.
func dummyString(format string) string {
	switch format {
	case "email":
		return "smoke-{{newUuid}}@example.com"
	case "date-time":
		return "2025-01-01T00:00:00Z"
	case "date":
		return "2025-01-01"
	case "uri", "url":
		return "https://example.com"
	case "password":
		return "testpassword123"
	default:
		return "test"
	}
}
