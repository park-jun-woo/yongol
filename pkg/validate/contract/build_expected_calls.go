//ff:func feature=validate-contract type=util control=sequence
//ff:what buildExpectedCalls — FuncSpec + SSaC.callRef 에서 허용 "pkg.Func" 호출 대상 집합 구축

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildExpectedCalls returns the set of `<pkg>.<Func>` identifiers a
// preserved handler may legally invoke. Sources:
//
//   - ProjectFuncSpecs / YongolPkgSpecs — canonical function surface.
//   - Ground "SSaC.callRef"              — normalized @call references.
//
// Both camelCase and PascalCase forms of each SSaC call are inserted
// so rendering variance between the parser and the user's code does
// not trip the check.
func buildExpectedCalls(fs *yongol.Fullstack, g *rule.Ground) map[string]bool {
	calls := map[string]bool{}
	appendFuncSpecEntries(fs.ProjectFuncSpecs, calls)
	appendFuncSpecEntries(fs.YongolPkgSpecs, calls)
	if g != nil {
		appendSSaCCallEntries(g.Lookup["SSaC.callRef"], calls)
	}
	return calls
}
