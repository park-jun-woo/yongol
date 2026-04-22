//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populatePathOpsParams — registers param/sort/filter from each operation in a path
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populatePathOpsParams(g *rule.Ground, ops map[string]*openapi3.Operation) {
	for _, op := range ops {
		// Operations without an operationId cannot produce a Lookup key — skip them.
		// operationId is optional per the OpenAPI spec, but yongol uses it as the
		// cross-SSOT join key, so the preceding validate step (O-4 rule in
		// pkg/validate/openapi) blocks these with an ERROR. The Ground layer is a
		// loader and never emits diagnostics; in practice this path is unreachable
		// after O-4 passes.
		if op.OperationID != "" {
			populateOpParams(g, op)
		}
	}
}
