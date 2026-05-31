//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectExternalPackages_ZeroCov(t *testing.T) {
	plansByFeature := map[string][]*ir.ServicePlan{
		"f": {{Ops: []ir.Op{
			{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "Charge"}},
			{Kind: ir.OpGet},
		}}},
	}
	got := collectExternalPackages(plansByFeature)
	if len(got) != 1 || got[0].Name != "billing" || got[0].Methods[0] != "Charge" {
		t.Errorf("unexpected: %+v", got)
	}
}
