//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what definedFeatureFuncs — feature 내 정의된 서비스 함수명 집합 구성 (snake_case)

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// definedFeatureFuncs returns the set of snake_case service function names
// defined by the given plans, used to skip stub generation for already-defined
// functions.
func definedFeatureFuncs(plans []*ir.ServicePlan) map[string]bool {
	defined := make(map[string]bool)
	for _, p := range plans {
		defined[snakeCase(p.OperationID)] = true
	}
	return defined
}
