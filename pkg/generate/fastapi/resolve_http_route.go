//ff:func feature=gen-fastapi type=util control=iteration dimension=2
//ff:what resolveHTTPRoute — OpenAPI operationId → HTTP method + URL path 해석

package fastapi

import (
	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveHTTPRoute looks up the OpenAPI doc to find the HTTP method and
// URL path for a given ServicePlan by matching its OperationID.
// FastAPI uses {param} syntax (same as OpenAPI), so no path conversion
// is needed.
func resolveHTTPRoute(plan *ir.ServicePlan, fs *yongol.Fullstack) {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return
	}
	for path, pi := range fs.OpenAPIDoc.Paths.Map() {
		for method, op := range pi.Operations() {
			if op != nil && op.OperationID == plan.OperationID {
				plan.HTTPMethod = method
				plan.URLPath = path // FastAPI uses OpenAPI {param} syntax natively
				return
			}
		}
	}
}
