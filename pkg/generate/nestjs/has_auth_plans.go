//ff:func feature=gen-nestjs type=util control=iteration dimension=2
//ff:what hasAuthPlans — feature별 ServicePlan 맵에 @auth Op 포함 여부 확인

package nestjs

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// hasAuthPlans returns true if any ServicePlan in the map contains an auth
// operation.
func hasAuthPlans(plansByFeature map[string][]*ir.ServicePlan) bool {
	for _, plans := range plansByFeature {
		for _, p := range plans {
			if containsOpKind(p.Ops, ir.OpAuth) {
				return true
			}
		}
	}
	return false
}
