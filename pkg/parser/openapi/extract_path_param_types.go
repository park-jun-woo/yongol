//ff:func feature=openapi-parse type=parser control=iteration dimension=2
//ff:what ExtractPathParamTypes — path 파라미터의 타입을 operationId별로 추출한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractPathParamTypes returns a map of operationId → paramName → OpenAPI type
// (e.g. "integer", "string"). Only path parameters are included.
func ExtractPathParamTypes(doc *openapi3.T) map[string]map[string]string {
	result := make(map[string]map[string]string)
	if doc == nil || doc.Paths == nil {
		return result
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range item.Operations() {
			collectPathParamTypesForOp(result, op)
		}
	}
	return result
}
