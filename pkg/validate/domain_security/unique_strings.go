//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what uniqueStrings — 문자열 슬라이스 중복 제거 (첫 발생 순서 유지)
package domain_security

// uniqueStrings returns deduplicated slice preserving first occurrence order.
func uniqueStrings(ss []string) []string {
	seen := make(map[string]struct{})
	var out []string
	for _, s := range ss {
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			out = append(out, s)
		}
	}
	return out
}
