//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what buildOperationMethodMap — OpenAPI operationId → (method, Operation) 맵 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// buildOperationMethodMap builds an operationId → (method, Operation) map
// from the OpenAPI doc.
func buildOperationMethodMap(doc *openapi3.T) map[string]operationEntry {
	out := make(map[string]operationEntry)
	if doc == nil || doc.Paths == nil {
		return out
	}
	for _, item := range doc.Paths.Map() {
		verbs := []struct {
			method string
			op     *openapi3.Operation
		}{
			{"GET", item.Get},
			{"POST", item.Post},
			{"PUT", item.Put},
			{"DELETE", item.Delete},
			{"PATCH", item.Patch},
		}
		for _, v := range verbs {
			if v.op == nil || v.op.OperationID == "" {
				continue
			}
			out[v.op.OperationID] = operationEntry{method: v.method, op: v.op}
		}
	}
	return out
}
