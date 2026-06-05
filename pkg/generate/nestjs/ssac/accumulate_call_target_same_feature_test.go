//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestAccumulateCallTargetSameFeature — TestAccumulateCallTarget — @call 타겟을 cross-feature/same-feature-stub/무관 op 로 분류 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAccumulateCallTargetSameFeature(t *testing.T) {
	deps := &moduleDeps{}
	callFeatures := map[string]bool{}
	op := ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{TargetFeature: "course"}}

	accumulateCallTarget(op, "course", deps, callFeatures)

	if !deps.NeedsSameFeatureStub {
		t.Error("NeedsSameFeatureStub should be true for same-feature call")
	}
	if len(callFeatures) != 0 {
		t.Errorf("callFeatures = %v, want empty for same-feature call", callFeatures)
	}
}
