//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestAccumulatePlanDepsNoSpecialOps — TestAccumulatePlanDeps — plan 의 publish/auth/@call 의존성을 deps·callFeatures 에 누적 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestAccumulatePlanDepsNoSpecialOps(t *testing.T) {
	p := &ir.ServicePlan{Ops: []ir.Op{{Kind: ir.OpGet}}}
	deps := &moduleDeps{}
	callFeatures := map[string]bool{}

	accumulatePlanDeps(p, "course", deps, callFeatures)

	if deps.NeedsQueue || deps.NeedsAuthz || len(callFeatures) != 0 {
		t.Errorf("expected no deps; deps=%+v callFeatures=%v", deps, callFeatures)
	}
}
