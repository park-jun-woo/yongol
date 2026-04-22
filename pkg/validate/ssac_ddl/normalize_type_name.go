//ff:func feature=validate type=util control=sequence topic=ssac-ddl
//ff:what normalizeTypeName — slice/pointer 접두사 제거

package ssac_ddl

// normalizeTypeName strips slice prefix and pointer prefix from a type name.
// e.g. "[]Reservation" → "Reservation", "*User" → "User"
func normalizeTypeName(t string) string {
	if len(t) > 2 && t[:2] == "[]" {
		t = t[2:]
	}
	if len(t) > 1 && t[0] == '*' {
		t = t[1:]
	}
	return t
}
