//ff:func feature=openapi-parse type=parser control=iteration dimension=2
//ff:what ExtractResponseArrayItemTypes — 응답 스키마에서 배열 필드의 항목 프로퍼티 타입을 추출한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractResponseArrayItemTypes returns a map of operationId → array field
// name → item property name → OpenAPI type (e.g. "integer", "string"). The
// react emitter consults it for row actions inside data-each: an item.<Field>
// mutate argument bound to an integer path parameter is wrapped with
// Number(...) only when the item field is not already a numeric type.
func ExtractResponseArrayItemTypes(doc *openapi3.T) map[string]map[string]map[string]string {
	result := make(map[string]map[string]map[string]string)
	if doc == nil || doc.Paths == nil {
		return result
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range item.Operations() {
			collectResponseArrayItemTypesForOp(result, op)
		}
	}
	return result
}
