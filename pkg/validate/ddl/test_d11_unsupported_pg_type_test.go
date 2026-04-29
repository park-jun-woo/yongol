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
			// DOUBLE PRECISION and TIMESTAMP WITH TIME ZONE are now
			// recognised via the alias matrix (FLOAT8 / TIMESTAMPTZ),
			// so they no longer fire D-11. The remaining unsupported
			// multi-word forms are TIME WITH/WITHOUT TIME ZONE and
			// BIT VARYING — none yet have a registered Go binding.
			name: "double precision now supported (zero diags)",
			columns: map[string]ddl.Column{
				"score":     {Name: "score", RawType: "DOUBLE PRECISION", NotNull: true},
				"occurred":  {Name: "occurred", RawType: "TIMESTAMP WITH TIME ZONE", NotNull: true},
				"naive":     {Name: "naive", RawType: "TIMESTAMP WITHOUT TIME ZONE", NotNull: true},
				"long_name": {Name: "long_name", RawType: "CHARACTER VARYING(255)", NotNull: true},
			},
			want: 0,
		},
		{
			name: "user-defined enum rejected",
			columns: map[string]ddl.Column{
				"status": {Name: "status", RawType: "ORDER_STATUS", NotNull: true},
			},
			want: 1,
		},
		{
			name: "time with tz / bit varying remain unsupported (two diags)",
			columns: map[string]ddl.Column{
				"clock": {Name: "clock", RawType: "TIME WITH TIME ZONE", NotNull: true},
				"flags": {Name: "flags", RawType: "BIT VARYING(8)", NotNull: true},
			},
			want: 2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) { runD11Case(t, c.columns, c.want) })
	}
}
