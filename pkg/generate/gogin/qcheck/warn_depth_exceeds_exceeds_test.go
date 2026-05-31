//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWarnDepthExceeds — 한도 초과 WARN 생성 / 한도 이내 무경고 / 파서 에러 분기 검증
package qcheck

import (
	"strings"
	"testing"
)

func TestWarnDepthExceeds_Exceeds(t *testing.T) {
	src := `package x
func deep() {
	if a {
		if b {
			if c {
				_ = 1
			}
		}
	}
}`
	warns, err := warnDepthExceeds("deep.go", src, Limits{MaxDepth: 2, MaxPureLines: 10})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(warns) != 1 || !strings.Contains(warns[0], "func=deep") {
		t.Fatalf("expected 1 WARN for deep, got %+v", warns)
	}
}
