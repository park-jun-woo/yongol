//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectOpsPackages_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{}
	ops := []ir.Op{
		{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Charge"}},
		{Kind: ir.OpGet},
	}
	collectOpsPackages(pm, ops)
	if !pm["billing"]["Charge"] {
		t.Error("expected billing.Charge collected")
	}
}
