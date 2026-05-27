//ff:func feature=gen-fastapi type=util control=iteration dimension=2
//ff:what collectImportData — ServicePlan 배열에서 import 에 필요한 데이터 수집

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// collectImportData scans all plans and returns the import metadata needed
// for writeServiceImports. currentFeature is used for self-import prevention.
func collectImportData(plans []*ir.ServicePlan, currentFeature string) importData {
	d := importData{
		Models:  make(map[string]bool),
		ExtPkgs: make(map[string]map[string]bool),
	}
	for _, plan := range plans {
		collectPlanImports(&d, plan.Ops, currentFeature)
	}
	return d
}
