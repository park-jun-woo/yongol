//ff:func feature=migration type=util control=sequence
//ff:what fkEqual — 두 ForeignKey 의 필드(RefTable/OnDelete/OnUpdate/Columns/RefColumns) 동등 비교
package migration

// fkEqual reports whether two FKs match on table/action/column lists.
func fkEqual(a, b *ForeignKey) bool {
	if a == nil || b == nil {
		return false
	}
	if a.RefTable != b.RefTable || a.OnDelete != b.OnDelete || a.OnUpdate != b.OnUpdate {
		return false
	}
	if !stringSliceEqual(a.Columns, b.Columns) {
		return false
	}
	if !stringSliceEqual(a.RefColumns, b.RefColumns) {
		return false
	}
	return true
}
