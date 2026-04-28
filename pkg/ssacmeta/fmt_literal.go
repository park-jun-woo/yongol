//ff:func feature=ssacmeta type=util control=sequence
//ff:what fmtLiteral — when: DSL equality 비교를 위한 문자열 캐스트

package ssacmeta

// fmtLiteral coerces v to its string literal form for equality comparisons
// in the when: DSL. Non-string values yield "" since the DSL only supports
// string literals on the right-hand side.
func fmtLiteral(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
