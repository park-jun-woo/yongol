//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what extractArrayItemTypes — 스키마의 배열 프로퍼티에서 항목 필드명→타입 맵 추출

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// extractArrayItemTypes returns a map of array field name → item property
// name → OpenAPI type from the given schema's properties.
func extractArrayItemTypes(schema *openapi3.Schema) map[string]map[string]string {
	fields := make(map[string]map[string]string)
	for propName, propRef := range schema.Properties {
		itemTypes := collectItemFieldTypes(propRef)
		if len(itemTypes) > 0 {
			fields[propName] = itemTypes
		}
	}
	return fields
}
