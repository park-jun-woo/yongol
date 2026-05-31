//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what collectSameFeatureStubs — 같은 feature 내 @call/@eval 대상 함수 목록 수집

package ssac

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// collectSameFeatureStubs scans all plans in a feature and returns the sorted
// list of @call/@eval function names whose package matches the current feature.
// These functions exist as inline calls but have no definition in the feature
// service file, so they need inline stubs appended at the bottom.
func collectSameFeatureStubs(plans []*ir.ServicePlan, feature string) []string {
	defined := definedFeatureFuncs(plans)

	stubs := make(map[string]bool)
	for _, p := range plans {
		collectPlanStubs(p, feature, defined, stubs)
	}

	result := make([]string, 0, len(stubs))
	for fn := range stubs {
		result = append(result, fn)
	}
	sort.Strings(result)
	return result
}
