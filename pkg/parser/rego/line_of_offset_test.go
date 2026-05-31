//ff:func feature=policy type=test control=iteration dimension=1
//ff:what TestRegoHelpers — unit tests for the pure rego parser helper functions
package rego

import (
	"testing"
)

func TestLineOfOffset(t *testing.T) {
	s := "line1\nline2\nline3"
	tests := []struct {
		off  int
		want int
	}{
		{0, 1},
		{5, 1},          // index of the first '\n' is on line 1
		{6, 2},          // first char of line2
		{12, 3},         // first char of line3
		{len(s), 3},     // end-of-string
		{-1, 0},         // out of range low
		{len(s) + 1, 0}, // out of range high
	}
	for _, tt := range tests {
		if got := lineOfOffset(s, tt.off); got != tt.want {
			t.Errorf("lineOfOffset(off=%d) = %d, want %d", tt.off, got, tt.want)
		}
	}
}
