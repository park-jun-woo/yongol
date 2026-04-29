//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what NormalizePGTypeHead — 다중 토큰 PG 타입 → 단일 토큰 별칭 회귀

package ddl

import "testing"

func TestNormalizePGTypeHead(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{in: "BIGINT", want: "BIGINT"},
		{in: "DOUBLE PRECISION", want: "FLOAT8"},
		{in: "TIMESTAMP WITH TIME ZONE", want: "TIMESTAMPTZ"},
		{in: "TIMESTAMP WITHOUT TIME ZONE", want: "TIMESTAMP"},
		{in: "TIME WITH TIME ZONE", want: "TIMETZ"},
		{in: "TIME WITHOUT TIME ZONE", want: "TIME"},
		{in: "CHARACTER VARYING", want: "VARCHAR"},
		{in: "CHARACTER", want: "CHAR"},
		{in: "BIT VARYING", want: "VARBIT"},
		{in: "  timestamp with time zone  ", want: "TIMESTAMPTZ"},
		{in: "INTERVAL", want: "INTERVAL"},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			got := NormalizePGTypeHead(c.in)
			if got != c.want {
				t.Errorf("NormalizePGTypeHead(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
