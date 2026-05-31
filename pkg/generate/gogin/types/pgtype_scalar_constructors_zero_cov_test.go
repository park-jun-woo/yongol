//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestPgtypeConstructorsZeroCov — 모든 pgtype 생성자 + unsupportedBinding 직접 커버
package types

import (
	"testing"
)

func TestPgtypeScalarConstructors_ZeroCov(t *testing.T) {
	cases := []struct {
		name string
		got  GoTypeBinding
		sqlc string
	}{
		{"bool", pgtypeBool("false"), "pgtype.Bool"},
		{"float4", pgtypeFloat4("0"), "pgtype.Float4"},
		{"float8", pgtypeFloat8("0"), "pgtype.Float8"},
		{"int8", pgtypeInt8("0"), "pgtype.Int8"},
		{"text", pgtypeText("''"), "pgtype.Text"},
	}
	for _, c := range cases {
		if c.got.SqlcGoType != c.sqlc {
			t.Errorf("%s SqlcGoType = %q, want %q", c.name, c.got.SqlcGoType, c.sqlc)
		}
		if c.got.Kind != KindPgtype || !c.got.Supported {
			t.Errorf("%s should be supported pgtype: %+v", c.name, c.got)
		}
	}
}
