//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestSplitTypeParams — base 와 괄호 내부 파라미터 분리
package migration

import (
	"testing"
)

func TestSplitTypeParams(t *testing.T) {
	cases := []struct {
		in         string
		wantBase   string
		wantParams string
	}{
		{"VARCHAR(255)", "VARCHAR", "255"},
		{"NUMERIC(10,2)", "NUMERIC", "10,2"},
		{"TEXT", "TEXT", ""},
		{"INTEGER[]", "INTEGER[]", ""},
	}
	for _, c := range cases {
		base, params := splitTypeParams(c.in)
		if base != c.wantBase || params != c.wantParams {
			t.Errorf("splitTypeParams(%q) = (%q,%q), want (%q,%q)", c.in, base, params, c.wantBase, c.wantParams)
		}
	}
}
