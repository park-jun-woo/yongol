//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what isCRUDSeq/parseDigits/splitModelMethod/planNeedsTransaction/csrfIsActive 순수 헬퍼
package ir

import (
	"testing"
)

func TestParseDigits(t *testing.T) {
	cases := []struct {
		in   string
		want int64
	}{
		{"123", 123},
		{"0", 0},
		{"", -1},
		{"12a", -1},
		{"1 2", -1},
		{"007", 7},
	}
	for _, c := range cases {
		if got := parseDigits(c.in); got != c.want {
			t.Errorf("parseDigits(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
