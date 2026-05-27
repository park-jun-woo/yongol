//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what collectExternalOpsPackages — Op 배열에서 @call/@eval 외부 패키지 수집

package ssac

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// collectExternalOpsPackages returns a sorted unique list of external packages
// referenced by @call/@eval ops within a single plan.
func collectExternalOpsPackages(ops []ir.Op) []string {
	seen := make(map[string]bool)
	for _, op := range ops {
		addExternalOpPackage(seen, op)
	}
	result := make([]string, 0, len(seen))
	for pkg := range seen {
		result = append(result, pkg)
	}
	sort.Strings(result)
	return result
}
