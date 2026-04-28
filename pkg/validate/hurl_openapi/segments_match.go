//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what segmentsMatch — 두 정규화 세그먼트 slice 의 element-wise 비교

package hurl_openapi

// segmentsMatch compares two normalized segment slices element-wise.
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
