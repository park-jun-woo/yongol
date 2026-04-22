//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what buildOperationMap — OpenAPI path 전체에서 operationId → Operation 맵 생성

package openapi_ssac

import "github.com/getkin/kin-openapi/openapi3"

// buildOperationMap builds an operationId -> Operation map from the OpenAPI doc.
func buildOperationMap(doc *openapi3.T) map[string]*openapi3.Operation {
	opMap := make(map[string]*openapi3.Operation)
	if doc == nil || doc.Paths == nil {
		return opMap
	}
	for _, item := range doc.Paths.Map() {
		addItemOperations(opMap, item)
	}
	return opMap
}
