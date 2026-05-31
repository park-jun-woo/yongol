//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what accumulatePlanDeps — 단일 plan 의 publish/auth/@call 의존성을 deps 와 callFeatures 에 누적

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// accumulatePlanDeps updates deps and callFeatures from a single plan's ops.
func accumulatePlanDeps(p *ir.ServicePlan, lowerFeature string, deps *moduleDeps, callFeatures map[string]bool) {
	if hasPublishOp(p.Ops) {
		deps.NeedsQueue = true
	}
	if hasAuthOp(p.Ops) {
		deps.NeedsAuthz = true
	}
	for _, op := range p.Ops {
		accumulateCallTarget(op, lowerFeature, deps, callFeatures)
	}
}
