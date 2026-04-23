//ff:func feature=validate type=rule control=iteration dimension=1 topic=tsx-openapi
//ff:what XOT-1 — verifies that the <op> in apiClient.<op>() calls exists as an OpenAPI operationId
package tsx_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xot01OperationID validates XOT-1: every apiClient.<op>() invocation in
// TSX must correspond to an OpenAPI operationId. Uses the shared Ground
// lookup built from fs.OpenAPIDoc.Paths so case / normalization matches
// the same set the other cross-validators compare against (XOS-15 etc.).
func xot01OperationID(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.TSXPages) == 0 {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	opIDs := g.Lookup["OpenAPI.operationId"]
	if len(opIDs) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, page := range fs.TSXPages {
		diags = append(diags, xot01CheckPage(page, opIDs)...)
	}
	return diags
}
