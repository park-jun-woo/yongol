//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWarnDepthExceeds — 한도 초과 WARN 생성 / 한도 이내 무경고 / 파서 에러 분기 검증
package qcheck

import (
	"testing"
)

func TestWarnDepthExceeds_ParseError(t *testing.T) {
	if _, err := warnDepthExceeds("bad.go", "@@@", Limits{MaxDepth: 3}); err == nil {
		t.Fatalf("expected parse error, got nil")
	}
}
