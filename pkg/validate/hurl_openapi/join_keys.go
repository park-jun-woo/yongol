//ff:func feature=validate type=util control=sequence topic=hurl-openapi
//ff:what joinKeys — 진단 메시지용 `[a, b, c]` 렌더링

package hurl_openapi

// joinKeys renders a key list as `[a, b, c]` for use in diagnostic
// messages.
func joinKeys(keys []string) string {
	if len(keys) == 0 {
		return "[]"
	}
	return "[" + joinCSV(keys) + "]"
}
