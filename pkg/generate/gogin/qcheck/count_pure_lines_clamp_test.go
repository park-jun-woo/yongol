//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCountPureLines — 공백/주석 제외 카운트 + 경계 clamp 검증
package qcheck

import (
	"testing"
)

func TestCountPureLines_Clamp(t *testing.T) {
	lines := []string{"x := 1", "y := 2"}
	// end far beyond len -> clamps to slice length; start=0.
	if got := countPureLines(lines, 0, 100); got != 2 {
		t.Errorf("clamped count = %d, want 2", got)
	}
	// Empty range.
	if got := countPureLines(lines, 0, 1); got != 0 {
		t.Errorf("empty range count = %d, want 0", got)
	}
}
