//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestMatchingCloseParen — 첫 "(" 의 짝 ")" 인덱스, 없으면 -1
package migration

import "testing"

func TestMatchingCloseParen(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"(abc)", 4},
		{"(a(b)c)", 6},
		{"(unbalanced", -1},
		{"()", 1},
	}
	for _, c := range cases {
		if got := matchingCloseParen(c.in); got != c.want {
			t.Errorf("matchingCloseParen(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
