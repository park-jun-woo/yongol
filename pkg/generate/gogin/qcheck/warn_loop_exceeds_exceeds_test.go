//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWarnLoopExceeds — 한도 초과 루프 WARN / 이내 무경고 / 파서 에러 분기 검증
package qcheck

import (
	"strings"
	"testing"
)

func TestWarnLoopExceeds_Exceeds(t *testing.T) {
	// longLoopSrc (fixtures_test.go) has a loop body with >10 pure lines.
	warns, err := warnLoopExceeds("long.go", longLoopSrc, Limits{MaxDepth: 3, MaxPureLines: 10})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(warns) != 1 || !strings.Contains(warns[0], "Q4 limit=10") {
		t.Fatalf("expected 1 Q4 WARN, got %+v", warns)
	}
}
