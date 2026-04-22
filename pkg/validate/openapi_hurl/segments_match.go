//ff:func feature=validate type=util control=iteration dimension=1 topic=scenario-check
//ff:what segmentsMatch — 두 정규화 세그먼트 배열의 일치 여부

package openapi_hurl

// segmentsMatch checks if two segment arrays match element-wise.
func segmentsMatch(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
