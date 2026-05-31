//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCanStripCastTail — 빈/괄호불균형/비식별자는 false, 단순 타입은 true
package migration

import (
	"testing"
)

func TestCanStripCastTail(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"text", true},
		{"", false},
		{"foo(", false}, // unbalanced parens
		{"a+b", false},  // not an identifier
		{"integer[]", true},
	}
	for _, c := range cases {
		if got := canStripCastTail(c.in); got != c.want {
			t.Errorf("canStripCastTail(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
