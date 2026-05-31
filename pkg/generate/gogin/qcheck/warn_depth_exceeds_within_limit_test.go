//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWarnDepthExceeds — 한도 초과 WARN 생성 / 한도 이내 무경고 / 파서 에러 분기 검증
package qcheck

import (
	"testing"
)

func TestWarnDepthExceeds_WithinLimit(t *testing.T) {
	src := "package x\nfunc shallow() {\n\tif a { _ = 1 }\n}"
	warns, err := warnDepthExceeds("ok.go", src, Limits{MaxDepth: 3, MaxPureLines: 10})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(warns) != 0 {
		t.Errorf("expected no warns within limit, got %+v", warns)
	}
}
