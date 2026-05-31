//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWarnLoopExceeds — 한도 초과 루프 WARN / 이내 무경고 / 파서 에러 분기 검증
package qcheck

import (
	"testing"
)

func TestWarnLoopExceeds_ParseError(t *testing.T) {
	if _, err := warnLoopExceeds("bad.go", "@@@", Limits{MaxPureLines: 10}); err == nil {
		t.Fatalf("expected parse error, got nil")
	}
}
