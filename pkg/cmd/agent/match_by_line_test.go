//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestMatchByLine — YAML line 에러의 라인 번호를 op 범위로 매핑 검증

package agent

import "testing"

func TestMatchByLine(t *testing.T) {
	offsets := []pathOffset{
		{Op: "OpA", StartLine: 1, EndLine: 10},
		{Op: "OpB", StartLine: 11, EndLine: 20},
	}
	got := matchByLine("yaml: line 15: bad indent", offsets)
	if len(got) != 1 || got[0] != "OpB" {
		t.Errorf("matchByLine line 15 = %v, want [OpB]", got)
	}
	// No line reference -> nil.
	if got := matchByLine("some unrelated error", offsets); got != nil {
		t.Errorf("expected nil for no line ref, got %v", got)
	}
	// Out-of-range line -> no op.
	if got := matchByLine("line 99 error", offsets); len(got) != 0 {
		t.Errorf("expected no op for out-of-range line, got %v", got)
	}
}
