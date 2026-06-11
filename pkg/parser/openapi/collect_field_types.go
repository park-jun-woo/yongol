//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what collectFieldTypes — 응답 스키마 top-level + allOf(1단계) 프로퍼티의 필드 경로별 타입 맵 생성

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectFieldTypes walks a response schema's top-level properties plus its
// allOf members (one level), expanding object and array properties into
// dotted-path keys.
func collectFieldTypes(schema *openapi3.Schema) map[string]FieldTypeInfo {
	out := make(map[string]FieldTypeInfo)
	addFieldTypes(out, schema)
	for _, allOf := range schema.AllOf {
		if allOf == nil || allOf.Value == nil {
			continue
		}
		addFieldTypes(out, allOf.Value)
	}
	return out
}
