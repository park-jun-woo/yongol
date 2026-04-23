//ff:func feature=validate type=rule control=iteration dimension=1 topic=tsx-openapi
//ff:what XOT-2 — verifies that path/query parameter object keys in apiClient calls exist in the OpenAPI parameters
package tsx_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xot02ParameterMatch validates XOT-2: for every apiClient.<op>({...})
// invocation, each top-level key of the argument object must match a path
// or query parameter declared on the OpenAPI operation.
//
// Keys that yongol / the generated api.ts uses as transport containers
// (e.g. `body`, `data`) are intentionally ignored so mutation-style calls
// (createWorkflow({ body: {...} })) don't produce false positives. Request
// body fields are the domain of XOT-3, not this rule.
//
// Skipped when XOT-1 would already fire for the same operation — since a
// missing operationId has no parameter set to compare against.
func xot02ParameterMatch(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.TSXPages) == 0 {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	opIDs := g.Lookup["OpenAPI.operationId"]
	var diags []diagnostic.Diagnostic
	for _, page := range fs.TSXPages {
		diags = append(diags, xot02CheckPage(page, opIDs, g.Lookup)...)
	}
	return diags
}
