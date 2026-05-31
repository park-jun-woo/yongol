//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasVerifyPasswordOp_ZeroCov(t *testing.T) {
	if !hasVerifyPasswordOp([]ir.Op{{Kind: ir.OpVerifyPassword}}) {
		t.Error("expected verify-password op")
	}
	if hasVerifyPasswordOp([]ir.Op{{Kind: ir.OpGet}}) {
		t.Error("unexpected verify-password op")
	}
}
