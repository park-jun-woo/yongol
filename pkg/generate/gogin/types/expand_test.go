//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestExpand — Expand 의 placeholder 치환 / escape / 누락 처리 회귀

package types

import "testing"

func TestExpand(t *testing.T) {
	cases := []struct {
		name string
		tmpl string
		row  string
		fld  string
		v    string
		want string
	}{
		{"empty template", "", "row", "F", "v", ""},
		{"all placeholders", "{row}.{field}={var}", "r", "F", "v", "r.F=v"},
		{"row only", "{row}.{field}", "row", "Foo", "", "row.Foo"},
		{"var only", "{var}", "", "", "src", "src"},
		{"escape lbrace", "{{x}}", "", "", "", "{x}"},
		{"missing placeholder substituted empty", "{row}.{field}", "", "F", "", ".F"},
		{"unused placeholder ignored", "{var}", "row", "F", "src", "src"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Expand(c.tmpl, c.row, c.fld, c.v); got != c.want {
				t.Errorf("Expand(%q,%q,%q,%q) = %q, want %q",
					c.tmpl, c.row, c.fld, c.v, got, c.want)
			}
		})
	}
}
