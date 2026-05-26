//ff:func feature=gen-typemap type=test control=iteration dimension=1
//ff:what ClassifyFamily — PGFamily 분류 단위 테스트 (14 family + Enum + Array + Unsupported)

package typemap

import "testing"

func TestClassifyFamily(t *testing.T) {
	cases := []struct {
		name string
		raw  string
		enum []string
		want PGFamily
	}{
		{"enum", "VARCHAR(50)", []string{"active", "inactive"}, FamilyEnum},
		{"enum overrides array", "VARCHAR(50)[]", []string{"a", "b"}, FamilyEnum},
		{"text array", "TEXT[]", nil, FamilyArray},
		{"uuid", "UUID", nil, FamilyUUID},
		{"numeric", "NUMERIC(10,2)", nil, FamilyNumeric},
		{"timestamptz", "TIMESTAMPTZ", nil, FamilyTimestampTZ},
		{"timestamp", "TIMESTAMP", nil, FamilyTimestamp},
		{"date", "DATE", nil, FamilyDate},
		{"inet", "INET", nil, FamilyInet},
		{"interval", "INTERVAL", nil, FamilyInterval},
		{"jsonb", "JSONB", nil, FamilyJSONB},
		{"bytea", "BYTEA", nil, FamilyBytea},
		{"bigint", "BIGINT", nil, FamilyInteger},
		{"serial", "SERIAL", nil, FamilyInteger},
		{"float8", "FLOAT8", nil, FamilyFloat},
		{"text", "TEXT", nil, FamilyString},
		{"boolean", "BOOLEAN", nil, FamilyBoolean},
		{"unsupported", "POINT", nil, FamilyUnsupported},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ClassifyFamily(simpleCol{raw: c.raw, enums: c.enum})
			if got != c.want {
				t.Errorf("got %s, want %s", got, c.want)
			}
		})
	}
}
