//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateResponseSchema — registers OpenAPI response fields per operationId into Ground.Schemas
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateResponseSchema(g *rule.Ground, opID string, op *openapi3.Operation) {
	if op.Responses == nil {
		return
	}
	primary2xxDone := false
	for code, resp := range op.Responses.Map() {
		primary2xxDone = applyResponseCodeSchema(g, opID, code, resp, primary2xxDone)
	}
}
