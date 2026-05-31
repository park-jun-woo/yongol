//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAddExternalOpPackage_ZeroCov(t *testing.T) {
	seen := map[string]bool{}
	addExternalOpPackage(seen, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing"}})
	addExternalOpPackage(seen, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "audit"}})
	if !seen["billing"] || !seen["audit"] {
		t.Errorf("packages not collected: %v", seen)
	}
}
