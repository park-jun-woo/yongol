//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectExternalOpsPackages_ZeroCov(t *testing.T) {
	ops := []ir.Op{
		{Kind: ir.OpCall, Call: &ir.CallOp{Package: "zeta"}},
		{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "alpha"}},
	}
	got := collectExternalOpsPackages(ops)
	if len(got) != 2 || got[0] != "alpha" || got[1] != "zeta" {
		t.Errorf("sorted packages wrong: %v", got)
	}
}
