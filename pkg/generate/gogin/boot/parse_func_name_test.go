//ff:func feature=gen-gogin type=test control=sequence
//ff:what parseFuncName — raw func 선언에서 식별자 이름만 추출

package boot

import "testing"

func TestParseFuncName(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"simple func", "func envInt(key string, def int) int { return 0 }", "envInt"},
		{"generic func", "func mapOf[T any](v T) {}", "mapOf"},
		{"no prefix", "var x = 1", ""},
		{"missing paren", "func broken", ""},
		{"empty", "", ""},
		{"immediate paren", "func ()", ""},
	}
	for _, c := range cases {
		if got := parseFuncName(c.in); got != c.want {
			t.Errorf("%s: parseFuncName(%q) = %q, want %q", c.name, c.in, got, c.want)
		}
	}
}
