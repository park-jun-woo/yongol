//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestAccumulateCallTargetNonCallOp — TestAccumulateCallTarget — @call 타겟을 cross-feature/same-feature-stub/무관 op 로 분류 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAccumulateCallTargetNonCallOp(t *testing.T) {
	deps := &moduleDeps{}
	callFeatures := map[string]bool{}

	// not a call op
	accumulateCallTarget(ir.Op{Kind: ir.OpGet}, "course", deps, callFeatures)
	// call op with empty TargetFeature
	accumulateCallTarget(ir.Op{Kind: ir.OpCall, Call: &ir.CallOp{}}, "course", deps, callFeatures)

	if deps.NeedsSameFeatureStub || len(callFeatures) != 0 {
		t.Errorf("expected no mutation; deps=%+v callFeatures=%v", deps, callFeatures)
	}
}
