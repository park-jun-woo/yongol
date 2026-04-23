//ff:func feature=migration type=util control=sequence
//ff:what riskyCast — VARCHAR 축소 / NUMERIC precision 축소 / 숫자↔텍스트 변환은 risky
package migration

// riskyCast reports whether altering a column from type `from` to `to`
// is risky enough to warrant MIG-005.
func riskyCast(from, to CanonicalType) bool {
	if from.Base == to.Base {
		return sameBaseShrink(from, to)
	}
	return crossCategoryCast(from.Base, to.Base)
}
