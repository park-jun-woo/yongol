//ff:func feature=stml-parse type=util control=sequence
//ff:what HTML 파서가 자동 생성하는 암시적 태그인지 판별
package stml

// isImplicitTag returns true for tags the HTML parser auto-generates.
func isImplicitTag(tag string) bool {
	return tag == "html" || tag == "head" || tag == "body"
}
