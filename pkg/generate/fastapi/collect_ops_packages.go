//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what collectOpsPackages — Op 배열에서 외부 패키지 참조 수집

package fastapi

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// collectOpsPackages adds external package references from ops to the map.
func collectOpsPackages(pm map[string]map[string]bool, ops []ir.Op) {
	for _, op := range ops {
		addOpPackageRef(pm, op)
	}
}
