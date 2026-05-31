//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCountPureLines — 공백/주석 제외 카운트 + 경계 clamp 검증
package qcheck

import (
	"testing"
)

func TestCountPureLines(t *testing.T) {
	lines := []string{
		"{",         // 0 (start)
		"a := 1",    // 1 pure
		"",          // 2 blank
		"  // note", // 3 comment
		"\tb := 2",  // 4 pure
		"}",         // 5
	}
	// start=1, end=6 -> counts lines[1..4] (end-1=5 exclusive): a:=1, blank, comment, b:=2 -> 2 pure.
	if got := countPureLines(lines, 1, 6); got != 2 {
		t.Errorf("countPureLines = %d, want 2", got)
	}
}
