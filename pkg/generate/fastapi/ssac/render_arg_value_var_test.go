//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderArgValueVar — renderArgValueVar LocVar source.col 표현식 렌더·폴백 분기 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderArgValueVar(t *testing.T) {
	cases := []struct {
		name string
		arg  ir.FieldArg
		col  string
		want string
	}{
		{
			name: "source column explicit",
			arg:  ir.FieldArg{Source: "order", SourceColumn: "org_id"},
			col:  "fallback",
			want: "order.org_id",
		},
		{
			name: "fallback to field name",
			arg:  ir.FieldArg{Source: "order", Field: ".OrgID"},
			col:  "fallback",
			want: "order.org_id",
		},
		{
			name: "fallback to column name",
			arg:  ir.FieldArg{Source: "order"},
			col:  "tenant_id",
			want: "order.tenant_id",
		},
		{
			name: "no source returns srcCol",
			arg:  ir.FieldArg{SourceColumn: "org_id"},
			col:  "fallback",
			want: "org_id",
		},
		{
			name: "source only, no column",
			arg:  ir.FieldArg{Source: "current_user"},
			col:  "",
			want: "current_user",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := renderArgValueVar(c.arg, c.col); got != c.want {
				t.Errorf("renderArgValueVar(%+v, %q) = %q, want %q", c.arg, c.col, got, c.want)
			}
		})
	}
}
