//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCanonicalType_SQL — VARCHAR(N)/CHAR(N)/NUMERIC(p,s)/array 렌더
package migration

import (
	"testing"
)

func TestCanonicalType_SQL(t *testing.T) {
	cases := []struct {
		name string
		t    CanonicalType
		want string
	}{
		{"plain", CanonicalType{Base: "INTEGER"}, "INTEGER"},
		{"varchar", CanonicalType{Base: "VARCHAR", Length: 255}, "VARCHAR(255)"},
		{"char", CanonicalType{Base: "CHAR", Length: 2}, "CHAR(2)"},
		{"numeric p", CanonicalType{Base: "NUMERIC", Precision: 10}, "NUMERIC(10)"},
		{"numeric ps", CanonicalType{Base: "NUMERIC", Precision: 10, Scale: 2}, "NUMERIC(10,2)"},
		{"array", CanonicalType{Base: "INTEGER", Array: true}, "INTEGER[]"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := c.t.SQL(); got != c.want {
				t.Errorf("SQL() = %q, want %q", got, c.want)
			}
		})
	}
}
