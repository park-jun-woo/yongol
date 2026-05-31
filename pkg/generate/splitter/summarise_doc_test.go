//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what snake / needsSnakeBreak / isVersionSuffix / tailSegment / summariseDoc 순수 헬퍼
package splitter

import (
	"testing"
)

func TestSummariseDoc(t *testing.T) {
	cases := []struct {
		name     string
		doc      string
		fallback string
		want     string
	}{
		{"first non-blank", "// first line\n// second", "fb", "first line"},
		{"skip leading blank", "\n\n// real", "fb", "real"},
		{"empty uses fallback", "", "MyFunc", "MyFunc"},
		{"empty no fallback", "   \n  ", "", "generated"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := summariseDoc(c.doc, c.fallback); got != c.want {
				t.Errorf("summariseDoc(%q,%q) = %q, want %q", c.doc, c.fallback, got, c.want)
			}
		})
	}
}
