//ff:func feature=gen-react type=helper control=iteration dimension=1
//ff:what findOpByID -- OpenAPI doc에서 operationId로 Operation을 찾는다

package react

import (
	"slices"

	"github.com/getkin/kin-openapi/openapi3"
)

func findOpByID(doc *openapi3.T, opID string) *openapi3.Operation {
	for _, item := range doc.Paths.Map() {
		ops := []*openapi3.Operation{item.Get, item.Post, item.Put, item.Delete, item.Patch}
		idx := slices.IndexFunc(ops, func(op *openapi3.Operation) bool { return op != nil && op.OperationID == opID })
		if idx >= 0 {
			return ops[idx]
		}
	}
	return nil
}
