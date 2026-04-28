//ff:func feature=crosscheck type=test-helper control=iteration dimension=1 topic=scenario-check
//ff:what containsAll — got 슬라이스가 want 요소 전부를 포함하는지 판정

package hurl

// containsAll reports whether got contains every element of want. Order
// is ignored; duplicates in want are safe (they hit the same set entry).
func containsAll(got, want []string) bool {
	set := map[string]struct{}{}
	for _, v := range got {
		set[v] = struct{}{}
	}
	for _, w := range want {
		if _, ok := set[w]; !ok {
			return false
		}
	}
	return true
}
