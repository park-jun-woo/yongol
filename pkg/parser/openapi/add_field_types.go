//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what addFieldTypes — 스키마 프로퍼티의 타입을 기록하고 object/array는 dotted 경로로 1단계 전개한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// addFieldTypes records each property's type and expands object/array
// properties into dotted "parent.child" keys one level deep.
func addFieldTypes(out map[string]FieldTypeInfo, schema *openapi3.Schema) {
	for name, ref := range schema.Properties {
		if ref == nil || ref.Value == nil {
			continue
		}
		out[name] = fieldTypeInfoOf(ref.Value)
		addNestedFieldTypes(out, name, ref.Value)
	}
}
