//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiHelpers — fastapi plan/package/route 헬퍼 검증 (Op 종류·외부 패키지 수집·라우트 해석)
package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAddOpPackageRefAndCollect(t *testing.T) {
	plans := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Hold"}},
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "", Function: "Local"}}, // empty pkg skipped
			{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "auth", Function: "IsExpired"}},
			{Kind: ir.OpEval, Eval: nil}, // nil eval skipped
			{Kind: ir.OpAuth},            // unhandled kind
		}}},
	}
	pkgs := collectExternalPackages(plans)
	if len(pkgs) != 2 {
		t.Fatalf("expected 2 packages, got %d: %+v", len(pkgs), pkgs)
	}
	// Sorted: auth before billing.
	if pkgs[0].Name != "auth" || pkgs[1].Name != "billing" {
		t.Errorf("unexpected package order: %+v", pkgs)
	}
	if len(pkgs[0].Methods) != 1 || pkgs[0].Methods[0] != "IsExpired" {
		t.Errorf("unexpected auth methods: %+v", pkgs[0])
	}
}
