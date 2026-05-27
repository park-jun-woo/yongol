//ff:func feature=gen-fastapi type=util control=iteration dimension=2
//ff:what hasPublishPlans — feature별 ServicePlan 맵에 @publish Op 포함 여부 확인

package fastapi

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// hasPublishPlans returns true if any ServicePlan in the map contains a
// publish operation.
func hasPublishPlans(plansByFeature map[string][]*ir.ServicePlan) bool {
	for _, plans := range plansByFeature {
		for _, p := range plans {
			if containsOpKind(p.Ops, ir.OpPublish) {
				return true
			}
		}
	}
	return false
}
