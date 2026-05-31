//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAddOpPackageRef_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{}
	addOpPackageRef(pm, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Charge"}})
	addOpPackageRef(pm, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "rules", Function: "IsExpired"}})
	// no-package call → ignored
	addOpPackageRef(pm, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: ""}})
	// other kind → ignored
	addOpPackageRef(pm, ir.Op{Kind: ir.OpGet})
	if !pm["billing"]["Charge"] {
		t.Error("expected billing.Charge")
	}
	if !pm["rules"]["IsExpired"] {
		t.Error("expected rules.IsExpired")
	}
}
