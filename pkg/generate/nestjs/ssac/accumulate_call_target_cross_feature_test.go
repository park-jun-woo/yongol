//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestAccumulateCallTargetCrossFeature — TestAccumulateCallTarget — @call 타겟을 cross-feature/same-feature-stub/무관 op 로 분류 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAccumulateCallTargetCrossFeature(t *testing.T) {
	deps := &moduleDeps{}
	callFeatures := map[string]bool{}
	op := ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{TargetFeature: "billing"}}

	accumulateCallTarget(op, "course", deps, callFeatures)

	if !callFeatures["billing"] {
		t.Errorf("callFeatures = %v, want billing", callFeatures)
	}
	if deps.NeedsSameFeatureStub {
		t.Error("NeedsSameFeatureStub should be false for cross-feature call")
	}
}
