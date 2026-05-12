//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what extractArrayItemFields — 스키마의 배열 프로퍼티에서 항목 프로퍼티 이름 맵 추출

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// extractArrayItemFields returns a map of array field name → set of item
// property names from the given schema's properties.
func extractArrayItemFields(schema *openapi3.Schema) map[string]map[string]bool {
	fields := make(map[string]map[string]bool)
	for propName, propRef := range schema.Properties {
		itemFields := collectItemFields(propRef)
		if len(itemFields) > 0 {
			fields[propName] = itemFields
		}
	}
	return fields
}
