//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what TestfmtLiteral — fmtLiteral() 문자열은 그대로, 비문자열은 "" 반환

package ssacmeta

import "testing"

func TestFmtLiteral(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want string
	}{
		{"string", "postgres", "postgres"},
		{"empty-string", "", ""},
		{"int", 5, ""},
		{"bool", true, ""},
		{"nil", nil, ""},
		{"float", 1.5, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := fmtLiteral(c.in); got != c.want {
				t.Errorf("fmtLiteral(%#v) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
