//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what collectModuleDeps — plan 들에서 queue/authz/same-feature-stub 필요 여부 + cross-feature @call 수집

package ssac

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// collectModuleDeps scans all plans for queue/authz needs and cross-feature
// (and same-feature) @call dependencies.
func collectModuleDeps(feature string, plans []*ir.ServicePlan) moduleDeps {
	lowerFeature := strings.ToLower(feature)
	var deps moduleDeps
	callFeatures := make(map[string]bool)

	for _, p := range plans {
		accumulatePlanDeps(p, lowerFeature, &deps, callFeatures)
	}

	deps.CrossFeatures = make([]string, 0, len(callFeatures))
	for cf := range callFeatures {
		deps.CrossFeatures = append(deps.CrossFeatures, cf)
	}
	sort.Strings(deps.CrossFeatures)
	return deps
}
