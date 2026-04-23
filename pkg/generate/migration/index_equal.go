//ff:func feature=migration type=util control=sequence
//ff:what indexEqual — 두 Index 의 Unique/Where/Columns 전체 동등 비교
package migration

// indexEqual reports whether two indexes are structurally identical.
func indexEqual(a, b *Index) bool {
	if a == nil || b == nil {
		return false
	}
	if a.Unique != b.Unique || a.Where != b.Where {
		return false
	}
	return stringSliceEqual(a.Columns, b.Columns)
}
