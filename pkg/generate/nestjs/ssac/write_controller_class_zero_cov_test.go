//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"strings"
	"testing"
)

func TestWriteControllerClass_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeControllerClass(&b, bnPlan())
	if !strings.Contains(b.String(), "@Controller(") {
		t.Errorf("controller class missing: %s", b.String())
	}
}
