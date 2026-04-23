//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-openapi
//ff:what buildOperationMethodMap — OpenAPI operationId → (method, Operation) 맵 생성

package openapi_ssac

import "github.com/getkin/kin-openapi/openapi3"

// buildOperationMethodMap builds an operationId → (method, Operation) map
// from the OpenAPI doc. buildOperationMap drops the method which is fine
// for the older rules but not for XOS-80/82 where the method determines
// the expected 2xx status.
func buildOperationMethodMap(doc *openapi3.T) map[string]OperationEntry {
	out := make(map[string]OperationEntry)
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
			out[v.op.OperationID] = OperationEntry{Method: v.method, Op: v.op}
		}
	}
	return out
}
