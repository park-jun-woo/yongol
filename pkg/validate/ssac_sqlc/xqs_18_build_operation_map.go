//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-sqlc
//ff:what buildXqs18OperationMap — build an operationId → *Operation map from the OpenAPI doc

package ssac_sqlc

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildXqs18OperationMap builds operationId → Operation map from the OpenAPI doc.
func buildXqs18OperationMap(doc *openapi3.T) map[string]*openapi3.Operation {
	opMap := make(map[string]*openapi3.Operation)
	if doc == nil || doc.Paths == nil {
		return opMap
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range []*openapi3.Operation{item.Get, item.Post, item.Put, item.Delete, item.Patch} {
			if op == nil || op.OperationID == "" {
				continue
			}
			opMap[op.OperationID] = op
		}
	}
	return opMap
}
