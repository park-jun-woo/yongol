//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what enumsMatch — 두 enum 슬라이스가 순서 무관하게 동일한지 비교

package openapi_ddl

import "sort"

// enumsMatch reports whether a and b contain the same elements (order-insensitive).
func enumsMatch(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sa := append([]string(nil), a...)
	sb := append([]string(nil), b...)
	sort.Strings(sa)
	sort.Strings(sb)
	for i := range sa {
		if sa[i] != sb[i] {
			return false
		}
	}
	return true
}
