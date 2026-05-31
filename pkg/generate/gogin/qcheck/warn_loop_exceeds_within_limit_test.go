//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWarnLoopExceeds — 한도 초과 루프 WARN / 이내 무경고 / 파서 에러 분기 검증
package qcheck

import (
	"testing"
)

func TestWarnLoopExceeds_WithinLimit(t *testing.T) {
	warns, err := warnLoopExceeds("long.go", longLoopSrc, Limits{MaxDepth: 3, MaxPureLines: 100})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(warns) != 0 {
		t.Errorf("expected no warns under high limit, got %+v", warns)
	}
}
