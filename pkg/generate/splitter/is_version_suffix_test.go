//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what snake / needsSnakeBreak / isVersionSuffix / tailSegment / summariseDoc 순수 헬퍼
package splitter

import (
	"testing"
)

func TestIsVersionSuffix(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"v2", true},
		{"v10", true},
		{"v", false},
		{"x2", false},
		{"v2a", false},
		{"", false},
	}
	for _, c := range cases {
		if got := isVersionSuffix(c.in); got != c.want {
			t.Errorf("isVersionSuffix(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
