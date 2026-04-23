//ff:func feature=migration type=parser control=sequence
//ff:what NormalizeType — DDL 타입 문자열을 CanonicalType 으로 변환 (aliases·SERIAL 해체)
package migration

import "strings"

// NormalizeType converts a raw DDL type string (e.g. "varchar(255)",
// "integer", "TIMESTAMP WITH TIME ZONE", "int4", "BIGSERIAL") into a
// CanonicalType. Unknown bases fall through with their uppercase literal.
//
// Returns (CanonicalType, isSerial). isSerial is true for SERIAL /
// BIGSERIAL / SMALLSERIAL so callers can attach a nextval() default.
func NormalizeType(raw string) (CanonicalType, bool) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return CanonicalType{}, false
	}
	array, s := stripTypeArraySuffix(s)
	base, params := splitTypeParams(s)
	upper := strings.ToUpper(strings.Join(strings.Fields(base), " "))

	ct := CanonicalType{Array: array}
	isSerial := normalizeTypeBase(upper, &ct)
	applyTypeParams(&ct, params)
	return ct, isSerial
}
