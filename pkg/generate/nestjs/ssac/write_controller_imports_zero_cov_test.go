//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"strings"
	"testing"
)

func TestWriteControllerImports_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeControllerImports(&b, bnPlan())
	out := b.String()
	if !strings.Contains(out, "Controller,") || !strings.Contains(out, "@nestjs/common") {
		t.Errorf("imports missing: %s", out)
	}
}
