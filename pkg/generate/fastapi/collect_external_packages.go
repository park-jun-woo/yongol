//ff:func feature=gen-fastapi type=util control=iteration dimension=2
//ff:what collectExternalPackages — ServicePlan 맵에서 @call/@eval 외부 패키지+메서드 수집

package fastapi

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// collectExternalPackages scans all plans for @call and @eval ops with a
// non-empty package and returns a sorted list of packages with their methods.
func collectExternalPackages(plansByFeature map[string][]*ir.ServicePlan) []externalPackage {
	pkgMethods := make(map[string]map[string]bool)
	for _, plans := range plansByFeature {
		for _, p := range plans {
			collectOpsPackages(pkgMethods, p.Ops)
		}
	}
	return buildSortedPackages(pkgMethods)
}
