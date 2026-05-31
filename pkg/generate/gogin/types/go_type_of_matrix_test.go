//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestGoTypeOf_Matrix(t *testing.T) {
	cases := []struct {
		raw  string
		want string
	}{
		{"BIGINT", "int64"},
		{"INTEGER", "int64"},
		{"SMALLINT", "int64"},
		{"SERIAL", "int64"},
		{"INT4", "int64"},
		{"VARCHAR(50)", "string"},
		{"TEXT", "string"},
		{"UUID", "string"},
		{"CHAR(2)", "string"},
		{"BOOLEAN", "bool"},
		{"BOOL", "bool"},
		{"TIMESTAMPTZ", "time.Time"},
		{"TIMESTAMP", "time.Time"},
		{"DATE", "time.Time"},
		{"NUMERIC(10,2)", "float64"},
		{"REAL", "float64"},
		{"FLOAT8", "float64"},
		{"JSONB", "json.RawMessage"},
		{"JSON", "json.RawMessage"},
		{"TEXT[]", "string"},      // array suffix stripped → TEXT
		{"BIGINT[]", "int64"},     // array suffix stripped → BIGINT
		{"MYSTERYTYPE", "string"}, // fallthrough default
	}
	for _, c := range cases {
		got := GoTypeOf(ddl.Column{RawType: c.raw})
		if got != c.want {
			t.Errorf("GoTypeOf(%q) = %q, want %q", c.raw, got, c.want)
		}
	}
}
