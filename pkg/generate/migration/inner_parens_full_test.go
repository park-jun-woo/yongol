//ff:func feature=migration type=test control=sequence
//ff:what TestInnerParensFull — 짝 맞는 ")" 까지 내부 반환 (중첩 지원)
package migration

import "testing"

func TestInnerParensFull(t *testing.T) {
	cases := []struct{ in, want string }{
		{"(a, b)", "a, b"},
		{"(f(x), g(y))", "f(x), g(y)"},
		{"no parens", "no parens"},
		{"(unterminated", "unterminated"},
	}
	for _, c := range cases {
		if got := innerParensFull(c.in); got != c.want {
			t.Errorf("innerParensFull(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
