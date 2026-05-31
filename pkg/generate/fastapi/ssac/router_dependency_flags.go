//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what routerDependencyFlags — plan 목록에서 인증/event_bus 의존성 필요 여부 판정

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// routerDependencyFlags reports whether any plan requires the authenticated
// user dependency and whether any plan publishes events.
func routerDependencyFlags(plans []*ir.ServicePlan) (needsAuth, needsEventBus bool) {
	for _, p := range plans {
		if p.TriggerKind == ir.TriggerHTTP && !hasVerifyPasswordOp(p.Ops) {
			needsAuth = true
		}
		if hasPublishOp(p.Ops) {
			needsEventBus = true
		}
	}
	return needsAuth, needsEventBus
}
