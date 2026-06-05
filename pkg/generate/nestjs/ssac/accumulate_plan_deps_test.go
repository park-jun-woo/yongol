//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestAccumulatePlanDeps — TestAccumulatePlanDeps — plan 의 publish/auth/@call 의존성을 deps·callFeatures 에 누적 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAccumulatePlanDeps(t *testing.T) {
	p := &ir.ServicePlan{
		Ops: []ir.Op{
			{Kind: ir.OpPublish},
			{Kind: ir.OpAuth},
			{Kind: ir.OpCall, Call: &ir.CallOp{TargetFeature: "billing"}},
		},
	}
	deps := &moduleDeps{}
	callFeatures := map[string]bool{}

	accumulatePlanDeps(p, "course", deps, callFeatures)

	if !deps.NeedsQueue {
		t.Error("NeedsQueue should be true (publish op)")
	}
	if !deps.NeedsAuthz {
		t.Error("NeedsAuthz should be true (auth op)")
	}
	if !callFeatures["billing"] {
		t.Errorf("callFeatures = %v, want billing", callFeatures)
	}
}
