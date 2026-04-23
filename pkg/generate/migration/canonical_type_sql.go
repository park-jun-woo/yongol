//ff:func feature=migration type=accessor control=selection
//ff:what CanonicalType.SQL — CanonicalType 를 PostgreSQL DDL 조각으로 렌더링
package migration

// SQL renders a CanonicalType back to PostgreSQL DDL fragment.
func (t CanonicalType) SQL() string {
	s := t.Base
	switch {
	case t.Length > 0 && t.Base == "VARCHAR":
		s = "VARCHAR(" + itoa(t.Length) + ")"
	case t.Length > 0 && t.Base == "CHAR":
		s = "CHAR(" + itoa(t.Length) + ")"
	case t.Precision > 0 && t.Base == "NUMERIC":
		if t.Scale > 0 {
			s = "NUMERIC(" + itoa(t.Precision) + "," + itoa(t.Scale) + ")"
		} else {
			s = "NUMERIC(" + itoa(t.Precision) + ")"
		}
	}
	if t.Array {
		s += "[]"
	}
	return s
}
