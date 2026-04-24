//ff:func feature=validate type=rule control=iteration dimension=1 topic=states
//ff:what xsm27DiagsForPath — 단일 OpenAPI path 에 대한 XSM-27 진단 수집

package ssac_statemachine

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// xsm27DiagsForPath yields XSM-27 diagnostics for POST/PUT/DELETE on a
// single OpenAPI path. Returns nil when the path is not stateful or has
// no `{id}` parameter. Extracted from xsm27StateIntentDeclaration so the
// outer walker stays at iteration dimension=1.
func xsm27DiagsForPath(
	pathStr string,
	item *openapi3.PathItem,
	diagrams []*statemachine.StateDiagram,
	g *rule.Ground,
	funcByName map[string]ssac.ServiceFunc,
) []diagnostic.Diagnostic {
	if item == nil {
		return nil
	}
	if !pathHasIDParam(pathStr) {
		return nil
	}
	target := isStatefulResource(pathStr, diagrams, g)
	if target == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, v := range statefulMethods(item) {
		if d, ok := xsm27DiagForOperation(v.method, v.op, target, funcByName); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
