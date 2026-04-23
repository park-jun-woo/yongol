//ff:func feature=migration type=util control=iteration dimension=1
//ff:what collectAllTableNames — prev/curr 테이블 이름 유니온 리스트 (중복 제거)
package migration

// collectAllTableNames returns the distinct table names present in
// either schema.
func collectAllTableNames(a, b *Schema) []string {
	set := map[string]bool{}
	for n := range a.Tables {
		set[n] = true
	}
	for n := range b.Tables {
		set[n] = true
	}
	out := make([]string, 0, len(set))
	for n := range set {
		out = append(out, n)
	}
	return out
}
