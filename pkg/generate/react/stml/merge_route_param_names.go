//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 두 파라미터 이름 슬라이스를 순서 유지하며 중복 없이 병합한다
package stml

// mergeRouteParamNames merges two param name slices, preserving order and
// removing duplicates.  Names from a come first, then unseen names from b.
func mergeRouteParamNames(a, b []string) []string {
	if len(a) == 0 && len(b) == 0 {
		return nil
	}
	seen := make(map[string]bool, len(a)+len(b))
	merged := make([]string, 0, len(a)+len(b))
	for _, n := range a {
		if !seen[n] {
			seen[n] = true
			merged = append(merged, n)
		}
	}
	for _, n := range b {
		if !seen[n] {
			seen[n] = true
			merged = append(merged, n)
		}
	}
	return merged
}
