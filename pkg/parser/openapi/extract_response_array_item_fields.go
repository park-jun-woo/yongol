//ff:func feature=openapi-parse type=parser control=iteration dimension=2
//ff:what ExtractResponseArrayItemFields — 응답 스키마에서 배열 필드의 항목 프로퍼티 이름을 추출한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractResponseArrayItemFields returns a map of operationId -> array field name
// -> set of item property names. This is used to determine whether list items
// have an "id" field for React key assignment.
func ExtractResponseArrayItemFields(doc *openapi3.T) map[string]map[string]map[string]bool {
	result := make(map[string]map[string]map[string]bool)
	if doc == nil || doc.Paths == nil {
		return result
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range item.Operations() {
			collectResponseArrayItemFieldsForOp(result, op)
		}
	}
	return result
}
