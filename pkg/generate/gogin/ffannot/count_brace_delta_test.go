//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCountBraceDelta — 라인 brace delta 계산 + 주석 제외 검증
package ffannot

import (
	"testing"
)

func TestCountBraceDelta(t *testing.T) {
	cases := map[string]int{
		"func f() {":            1,
		"}":                     -1,
		"x := map[string]int{}": 0,
		"{{}":                   1,
		"} // closing } brace":  -1, // braces in comment ignored
		"// { { {":              0,  // entire line is a comment
		"no braces here":        0,
	}
	for line, want := range cases {
		if got := countBraceDelta(line); got != want {
			t.Errorf("countBraceDelta(%q) = %d, want %d", line, got, want)
		}
	}
}
