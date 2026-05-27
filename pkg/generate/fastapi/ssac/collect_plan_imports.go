//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what collectPlanImports — 단일 plan 의 Ops → importData 수집

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// collectPlanImports processes a single plan's ops into the importData.
// currentFeature is forwarded to collectOpImport for self-import prevention.
func collectPlanImports(d *importData, ops []ir.Op, currentFeature string) {
	for _, op := range ops {
		collectOpImport(d, op, currentFeature)
	}
}
