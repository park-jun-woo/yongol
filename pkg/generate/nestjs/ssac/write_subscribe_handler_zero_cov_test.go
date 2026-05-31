//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"strings"
	"testing"
)

func TestWriteSubscribeHandler_ZeroCov(t *testing.T) {
	var b strings.Builder
	p := bnPlan()
	p.Topic = "order.completed"
	writeSubscribeHandler(&b, p)
	if !strings.Contains(b.String(), "order.completed") {
		t.Errorf("subscribe handler missing: %s", b.String())
	}
}
