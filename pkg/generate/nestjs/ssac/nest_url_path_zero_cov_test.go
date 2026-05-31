//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"
)

func TestNestURLPath_ZeroCov(t *testing.T) {
	if got := nestURLPath("/orders/{id}/items/{itemId}"); got != "/orders/:id/items/:itemId" {
		t.Errorf("nestURLPath=%q", got)
	}
	if got := nestURLPath("/plain"); got != "/plain" {
		t.Errorf("plain path changed: %q", got)
	}
}
