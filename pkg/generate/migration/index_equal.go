//ff:func feature=migration type=util control=sequence
//ff:what indexEqual — 두 Index 의 Unique/Where/Columns 전체 동등 비교
package migration

// indexEqual reports whether two indexes are structurally identical.
// Method treats "" and "btree" as equivalent since PostgreSQL's default
// access method is btree; the emitter omits `USING btree` for brevity,
// so the snapshot/baseline round-trip never loses equality.
func indexEqual(a, b *Index) bool {
	if a == nil || b == nil {
		return false
	}
	if a.Unique != b.Unique || a.Where != b.Where {
		return false
	}
	if normalizeIndexMethod(a.Method) != normalizeIndexMethod(b.Method) {
		return false
	}
	return stringSliceEqual(a.Columns, b.Columns)
}

// normalizeIndexMethod canonicalises empty string and "btree" to the same
// token so equality treats them interchangeably.
func normalizeIndexMethod(m string) string {
	if m == "" {
		return "btree"
	}
	return m
}
