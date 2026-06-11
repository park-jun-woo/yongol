//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what addChildFieldTypes — 하위 프로퍼티 집합을 parent.child 키로 타입 맵에 추가한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// addChildFieldTypes records each child property's type under the
// "parent.child" dotted key.
func addChildFieldTypes(out map[string]FieldTypeInfo, parent string, props openapi3.Schemas) {
	for sub, ref := range props {
		if ref == nil || ref.Value == nil {
			continue
		}
		out[parent+"."+sub] = fieldTypeInfoOf(ref.Value)
	}
}
