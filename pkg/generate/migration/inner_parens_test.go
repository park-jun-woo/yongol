//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestInnerParens — 바깥 "()" 제거, 비괄호는 그대로
package migration

import (
	"testing"
)

func TestInnerParens(t *testing.T) {
	cases := []struct{ in, want string }{
		{"(age >= 0)", "age >= 0"},
		{"  ( x ) ", "x"},
		{"no parens", "no parens"},
	}
	for _, c := range cases {
		if got := innerParens(c.in); got != c.want {
			t.Errorf("innerParens(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
