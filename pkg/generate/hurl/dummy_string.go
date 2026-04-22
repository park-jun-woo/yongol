//ff:func feature=gen-hurl type=util control=selection
//ff:what dummyString — string format에 따른 dummy 문자열 값 생성
package hurl

// dummyString returns a dummy string value based on the format.
func dummyString(format string) string {
	switch format {
	case "email":
		return "test@example.com"
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
