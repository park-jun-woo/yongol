//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestAddOpPackageRef_ZeroCov — Op(Call/Eval)에서 외부 패키지 참조 수집 분기 직접 호출

package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAddOpPackageRef_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{}

	// OpCall with package ref.
	addOpPackageRef(pm, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "HoldEscrow"}})
	if !pm["billing"]["HoldEscrow"] {
		t.Errorf("call package ref not added: %v", pm)
	}

	// OpEval with package ref.
	addOpPackageRef(pm, ir.Op{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "policy", Function: "IsExpired"}})
	if !pm["policy"]["IsExpired"] {
		t.Errorf("eval package ref not added: %v", pm)
	}

	// OpCall with empty package → no-op.
	addOpPackageRef(pm, ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{Function: "Local"}})
	// OpCall nil Call pointer → no-op.
	addOpPackageRef(pm, ir.Op{Kind: ir.OpCall})
	// Unrelated kind → no-op.
	addOpPackageRef(pm, ir.Op{Kind: ir.OpGet})
	if len(pm) != 2 {
		t.Errorf("unexpected package map growth: %v", pm)
	}
}
