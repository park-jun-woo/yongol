//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what isOverridden — file+class 조합이 override 셋에 있는지 확인
package stml_design

// isOverridden returns true if the given file+class combination is in the override set.
func isOverridden(ovr overrideSet, file, class string) bool {
	if class == "" {
		return false
	}
	m, ok := ovr[file]
	if !ok {
		return false
	}
	return m[class]
}
