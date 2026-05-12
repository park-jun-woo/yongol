//ff:func feature=openapi-parse type=parser control=iteration dimension=2
//ff:what ExtractNoBodyOps — requestBody가 없는 operationId 집합을 반환한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractNoBodyOps returns a set of operationIds whose operations have no
// requestBody defined. These operations produce void mutations in React
// (mutate() instead of mutate({})).
func ExtractNoBodyOps(doc *openapi3.T) map[string]bool {
	result := make(map[string]bool)
	if doc == nil || doc.Paths == nil {
		return result
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range item.Operations() {
			if op.OperationID != "" && op.RequestBody == nil {
				result[op.OperationID] = true
			}
		}
	}
	return result
}
