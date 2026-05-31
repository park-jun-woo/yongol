//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what collectPlanStubs — 단일 plan 의 @call/@eval op 중 미정의 same-feature 함수를 stubs 에 수집

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// collectPlanStubs records, into stubs, the snake_case function names of
// @call/@eval ops in plan that target the given feature and are not already
// defined.
func collectPlanStubs(plan *ir.ServicePlan, feature string, defined, stubs map[string]bool) {
	for _, op := range plan.Ops {
		pkg, fn := callEvalTarget(op)
		if pkg != feature || fn == "" {
			continue
		}
		snake := snakeCase(fn)
		if !defined[snake] {
			stubs[snake] = true
		}
	}
}
