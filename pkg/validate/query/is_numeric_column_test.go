//ff:func feature=validate type=test control=iteration dimension=1 topic=query-structural
//ff:what isNumericColumn — NUMERIC/DECIMAL 타입 판정 (매칭/파라미터/비매칭) 검증

package query

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsNumericColumn(t *testing.T) {
	cases := []struct {
		name    string
		rawType string
		want    bool
	}{
		{name: "NUMERIC", rawType: "NUMERIC", want: true},
		{name: "numeric_lower", rawType: "numeric", want: true},
		{name: "NUMERIC_with_precision", rawType: "NUMERIC(10,2)", want: true},
		{name: "DECIMAL", rawType: "DECIMAL", want: true},
		{name: "decimal_lower", rawType: "decimal", want: true},
		{name: "DECIMAL_with_precision", rawType: "DECIMAL(18,4)", want: true},
		{name: "numeric_array", rawType: "NUMERIC[]", want: true},
		{name: "not_numeric_bigint", rawType: "BIGINT", want: false},
		{name: "not_numeric_integer", rawType: "INTEGER", want: false},
		{name: "empty", rawType: "", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			col := ddl.Column{RawType: c.rawType}
			got := isNumericColumn(col)
			if got != c.want {
				t.Errorf("isNumericColumn(%q) = %v, want %v", c.rawType, got, c.want)
			}
		})
	}
}
