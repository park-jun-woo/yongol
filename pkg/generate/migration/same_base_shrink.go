//ff:func feature=migration type=util control=sequence
//ff:what sameBaseShrink — 같은 Base 인데 VARCHAR/NUMERIC 길이·정밀도 축소인지 판정
package migration

// sameBaseShrink returns true when from/to share a Base but to is
// narrower (VARCHAR length shrink, NUMERIC precision shrink).
func sameBaseShrink(from, to CanonicalType) bool {
	if from.Base == "VARCHAR" && from.Length > 0 && to.Length > 0 && to.Length < from.Length {
		return true
	}
	if from.Base == "NUMERIC" && from.Precision > 0 && to.Precision > 0 && to.Precision < from.Precision {
		return true
	}
	return false
}
