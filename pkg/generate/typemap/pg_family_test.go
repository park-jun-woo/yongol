//ff:func feature=gen-typemap type=test control=iteration dimension=1
//ff:what PGFamily.String — PGFamily 열거형의 문자열 표현 검증

package typemap

import "testing"

func TestPGFamilyString(t *testing.T) {
	cases := []struct {
		f    PGFamily
		want string
	}{
		{FamilyInteger, "Integer"},
		{FamilyFloat, "Float"},
		{FamilyString, "String"},
		{FamilyBoolean, "Boolean"},
		{FamilyUUID, "UUID"},
		{FamilyNumeric, "Numeric"},
		{FamilyTimestamp, "Timestamp"},
		{FamilyTimestampTZ, "TimestampTZ"},
		{FamilyDate, "Date"},
		{FamilyInet, "Inet"},
		{FamilyInterval, "Interval"},
		{FamilyJSONB, "JSONB"},
		{FamilyBytea, "Bytea"},
		{FamilyEnum, "Enum"},
		{FamilyArray, "Array"},
		{FamilyUnsupported, "Unsupported"},
		{PGFamily(999), "Unknown"},
	}
	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			got := c.f.String()
			if got != c.want {
				t.Errorf("PGFamily(%d).String() = %q, want %q", c.f, got, c.want)
			}
		})
	}
}
