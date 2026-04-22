//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what containsField — 필드 목록에 주어진 이름이 있는지 판정

package ssac

// containsField reports whether f exists in the fields slice.
func containsField(fields []string, f string) bool {
	for _, x := range fields {
		if x == f {
			return true
		}
	}
	return false
}
