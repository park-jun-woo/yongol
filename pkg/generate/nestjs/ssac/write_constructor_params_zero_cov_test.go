//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"strings"
	"testing"
)

func TestWriteConstructorParams_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeConstructorParams(&b, bnPlan())
	out := b.String()
	if !strings.Contains(out, "constructor(") || !strings.Contains(out, "QueueService") || !strings.Contains(out, "AuthzService") {
		t.Errorf("constructor params missing pieces: %s", out)
	}
}
