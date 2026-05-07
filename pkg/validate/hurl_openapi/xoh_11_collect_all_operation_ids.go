//ff:func feature=validate type=util control=iteration dimension=2 topic=hurl-openapi
//ff:what collectAllOperationIDs — OpenAPI Doc 의 전체 operationId 집합 수집

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

func collectAllOperationIDs(doc *openapi3.T) map[string]bool {
	ids := map[string]bool{}
	if doc == nil || doc.Paths == nil {
		return ids
	}
	for _, pi := range doc.Paths.Map() {
		addOpsFromPathItem(ids, pi)
	}
	return ids
}
