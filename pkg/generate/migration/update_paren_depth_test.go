//ff:func feature=migration type=test control=selection
//ff:what TestUpdateParenDepth — '(' +1, ')' -1, 그 외 유지
package migration

import "testing"

func TestUpdateParenDepth(t *testing.T) {
	cases := []struct {
		c     byte
		depth int
		want  int
	}{
		{'(', 0, 1},
		{')', 2, 1},
		{'x', 3, 3},
	}
	for _, c := range cases {
		if got := updateParenDepth(c.c, c.depth); got != c.want {
			t.Errorf("updateParenDepth(%q,%d) = %d, want %d", c.c, c.depth, got, c.want)
		}
	}
}
