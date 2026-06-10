//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what opSecurityIndex — OpenAPI 문서의 operationId → security 보호 여부 인덱스 구축

package react

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/validate/stml_openapi"
)

// opSecurityIndex maps every non-empty operationId in the OpenAPI document
// to whether it requires auth, using the same OpenAPI 3 security inheritance
// judgment the validator applies (stml_openapi.OpRequiresAuth — op.Security
// wins, nil inherits doc.Security, explicit [] opts out). The key set doubles
// as the operationId universe for component api.<Op>( consumption filtering.
func opSecurityIndex(doc *openapi3.T) map[string]bool {
	out := make(map[string]bool)
	if doc == nil || doc.Paths == nil {
		return out
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range item.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			out[op.OperationID] = stml_openapi.OpRequiresAuth(op, doc)
		}
	}
	return out
}
