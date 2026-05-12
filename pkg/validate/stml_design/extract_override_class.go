//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what extractOverrideClass — @override 주석에서 class 값 추출
package stml_design

// extractOverrideClass extracts the class value from an @override comment.
// Returns "" if no class attribute is present (structure-only override).
func extractOverrideClass(data string) string {
	m := overrideClassRe.FindStringSubmatch(data)
	if m == nil {
		return ""
	}
	return m[1]
}
