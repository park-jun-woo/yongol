//ff:func feature=gen-fastapi type=util control=iteration dimension=2
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
	// Collect all service function names defined in this feature to avoid
	// generating stubs for functions that already have definitions.
	defined := make(map[string]bool)
	for _, p := range plans {
		defined[snakeCase(p.OperationID)] = true
	}

	stubs := make(map[string]bool)
	for _, p := range plans {
		for _, op := range p.Ops {
			switch op.Kind {
			case ir.OpCall:
				if op.Call != nil && op.Call.Package == feature {
					fn := snakeCase(op.Call.Function)
					if !defined[fn] {
						stubs[fn] = true
					}
				}
			case ir.OpEval:
				if op.Eval != nil && op.Eval.Package == feature {
					fn := snakeCase(op.Eval.Function)
					if !defined[fn] {
						stubs[fn] = true
					}
				}
			}
		}
	}

	result := make([]string, 0, len(stubs))
	for fn := range stubs {
		result = append(result, fn)
	}
	sort.Strings(result)
	return result
}
