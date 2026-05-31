//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestFastapiPackages_ZeroCov — addOpPackageRef/collectOpsPackages/buildSortedPackages 커버

package fastapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAddOpPackageRefAndCollect_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{}
	ops := []ir.Op{
		{Kind: ir.OpCall, Call: &ir.CallOp{Package: "mail", Function: "Send"}},
		{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "rule", Function: "Check"}},
		{Kind: ir.OpCall, Call: &ir.CallOp{Package: ""}},   // empty pkg → skipped
		{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: ""}},   // empty pkg → skipped
		{Kind: ir.OpCall},                                  // nil Call → skipped
		{Kind: ir.OpEval},                                  // nil Eval → skipped
		{Kind: ir.OpGet, Get: &ir.GetOp{}},                 // other kind → no-op
	}
	collectOpsPackages(pm, ops)
	if !pm["mail"]["Send"] || !pm["rule"]["Check"] {
		t.Fatalf("expected mail.Send and rule.Check, got %v", pm)
	}
	if _, ok := pm[""]; ok {
		t.Error("empty package should not be added")
	}
}

func TestBuildSortedPackages_ZeroCov(t *testing.T) {
	pm := map[string]map[string]bool{
		"zeta":  {"B": true, "A": true},
		"alpha": {"X": true},
	}
	got := buildSortedPackages(pm)
	if len(got) != 2 || got[0].Name != "alpha" || got[1].Name != "zeta" {
		t.Fatalf("packages not sorted: %+v", got)
	}
	if len(got[1].Methods) != 2 || got[1].Methods[0] != "A" || got[1].Methods[1] != "B" {
		t.Errorf("methods not sorted: %+v", got[1].Methods)
	}
}
