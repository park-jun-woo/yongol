//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestMaxBlockDepth — for/range/switch/typeswitch/select/case 각 노드 종류별 depth 검증
package qcheck

import (
	"testing"
)

func TestMaxBlockDepth_Kinds(t *testing.T) {
	cases := []struct {
		name string
		body string
		want int
	}{
		{"flat", "_ = 1", 0},
		{"for", "for i := 0; i < 1; i++ { _ = i }", 1},
		{"range", "for _, v := range []int{1} { _ = v }", 1},
		{"switch", "switch { case true: _ = 1 }", 1},
		{"typeswitch", "var a any; switch a.(type) { case int: _ = 1 }", 1},
		{"select", "select { default: _ = 1 }", 1},
		{"nestedForInIf", "if c() { for { break } }", 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := maxBlockDepth(bodyBlock(t, c.body), 0)
			if got != c.want {
				t.Errorf("maxBlockDepth(%q) = %d, want %d", c.body, got, c.want)
			}
		})
	}
}
