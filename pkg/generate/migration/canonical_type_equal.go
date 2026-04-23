//ff:func feature=migration type=accessor control=sequence
//ff:what CanonicalType.Equal — 두 CanonicalType 값 전체 필드 비교
package migration

// Equal reports whether two CanonicalType values are identical in all
// fields. Helper for diff engines.
func (t CanonicalType) Equal(other CanonicalType) bool {
	return t.Base == other.Base &&
		t.Length == other.Length &&
		t.Precision == other.Precision &&
		t.Scale == other.Scale &&
		t.Array == other.Array
}
