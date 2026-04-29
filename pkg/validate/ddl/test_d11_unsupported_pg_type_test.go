//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestD11UnsupportedPgType — 다중 토큰 / 미인식 PG 타입 컬럼이 D-11 ERROR 를 발생시키는지

package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestD11UnsupportedPgType(t *testing.T) {
	cases := []struct {
		name    string
		columns map[string]ddl.Column
		want    int
	}{
		{
			name: "all supported (zero diags)",
			columns: map[string]ddl.Column{
				"id":   {Name: "id", RawType: "BIGINT", NotNull: true},
				"name": {Name: "name", RawType: "TEXT", NotNull: true},
				"uuid": {Name: "uuid", RawType: "UUID", NotNull: true},
			},
			want: 0,
		},
		{
			name: "double precision rejected",
			columns: map[string]ddl.Column{
				"score": {Name: "score", RawType: "DOUBLE PRECISION", NotNull: true},
			},
			want: 1,
		},
		{
			name: "user-defined enum rejected",
			columns: map[string]ddl.Column{
				"status": {Name: "status", RawType: "ORDER_STATUS", NotNull: true},
			},
			want: 1,
		},
		{
			name: "two unsupported produce two diags",
			columns: map[string]ddl.Column{
				"a": {Name: "a", RawType: "DOUBLE PRECISION", NotNull: true},
				"b": {Name: "b", RawType: "TIMESTAMP WITH TIME ZONE", NotNull: true},
			},
			want: 2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) { runD11Case(t, c.columns, c.want) })
	}
}
