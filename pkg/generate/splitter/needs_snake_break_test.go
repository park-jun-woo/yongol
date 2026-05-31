//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what snake / needsSnakeBreak / isVersionSuffix / tailSegment / summariseDoc 순수 헬퍼
package splitter

import (
	"testing"
)

func TestNeedsSnakeBreak(t *testing.T) {
	cases := []struct {
		s    string
		i    int
		want bool
	}{
		{"CamelCase", 0, false},  // i==0
		{"CamelCase", 5, true},   // l->C boundary (Camel|Case)
		{"HTTPServer", 4, true},  // acronym end: P->S where S precedes lowercase e
		{"HTTPServer", 1, false}, // T after H, both upper, next upper -> no break
		{"abc", 1, false},        // not upper
		{"x9Y", 2, true},         // digit before upper
	}
	for _, c := range cases {
		if got := needsSnakeBreak([]rune(c.s), c.i); got != c.want {
			t.Errorf("needsSnakeBreak(%q,%d) = %v, want %v", c.s, c.i, got, c.want)
		}
	}
}
