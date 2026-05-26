//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what buildPlansByFeature — SSaC ServiceFunc 배열 → feature별 ServicePlan 맵 변환

package fastapi

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildPlansByFeature iterates all SSaC service functions, builds a
// ServicePlan for each, resolves HTTP routes, and groups by feature name.
func buildPlansByFeature(fs *yongol.Fullstack) map[string][]*ir.ServicePlan {
	result := make(map[string][]*ir.ServicePlan)
	for i := range fs.ServiceFuncs {
		sf := &fs.ServiceFuncs[i]
		plan, err := ir.BuildServicePlan(sf, fs)
		if err != nil {
			continue // skip malformed entries; validate catches them earlier
		}
		if plan.TriggerKind == ir.TriggerHTTP {
			resolveHTTPRoute(plan, fs)
		}
		result[plan.Feature] = append(result[plan.Feature], plan)
	}
	return result
}
