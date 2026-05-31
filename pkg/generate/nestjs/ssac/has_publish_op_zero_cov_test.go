//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasPublishOp_ZeroCov(t *testing.T) {
	if !hasPublishOp([]ir.Op{{Kind: ir.OpPublish}}) {
		t.Error("expected publish op")
	}
	if hasPublishOp([]ir.Op{{Kind: ir.OpGet}}) {
		t.Error("unexpected publish op")
	}
}
