//ff:func feature=gen-gogin type=test control=iteration topic=loop-report
//ff:what TestCountPureLines — 공백/주석 제외 카운트 + 경계 clamp 검증

package qcheck

import "testing"

func TestCountPureLines(t *testing.T) {
	lines := []string{
		"{",          // 0 (start)
		"a := 1",     // 1 pure
		"",           // 2 blank
		"  // note",  // 3 comment
		"\tb := 2",   // 4 pure
		"}",          // 5
	}
	// start=1, end=6 -> counts lines[1..4] (end-1=5 exclusive): a:=1, blank, comment, b:=2 -> 2 pure.
	if got := countPureLines(lines, 1, 6); got != 2 {
		t.Errorf("countPureLines = %d, want 2", got)
	}
}

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
