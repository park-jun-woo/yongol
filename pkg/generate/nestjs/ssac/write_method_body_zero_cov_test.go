//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteMethodBody_ZeroCov(t *testing.T) {
	var b strings.Builder
	p := bnPlan()
	p.UsesTransaction = true
	writeMethodBody(&b, p)
	if !strings.Contains(b.String(), "$transaction") {
		t.Errorf("transaction body missing: %s", b.String())
	}
	var b2 strings.Builder
	sub := bnPlan()
	sub.TriggerKind = ir.TriggerSubscribe
	writeMethodBody(&b2, sub)
	if !strings.Contains(b2.String(), "const message = payload") {
		t.Errorf("subscribe alias missing: %s", b2.String())
	}
}
