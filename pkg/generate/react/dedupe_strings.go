//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what dedupeStrings — 순서 보존 문자열 중복 제거 (메뉴 조상 하이라이트 prefix 정리용)

package react

// dedupeStrings removes duplicates from a string slice preserving the
// first-occurrence order — the ancestor-highlight prefixes of a menu item
// stay in document order without repeating a startsWith condition.
func dedupeStrings(in []string) []string {
	seen := make(map[string]struct{}, len(in))
	var out []string
	for _, s := range in {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		out = append(out, s)
	}
	return out
}
