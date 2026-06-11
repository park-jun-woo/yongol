//ff:func feature=openapi-parse type=parser control=sequence
//ff:what addNestedFieldTypes — object/array 프로퍼티의 하위 필드를 parent.child 키로 1단계 전개한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// addNestedFieldTypes expands an object property's properties (or an array
// property's item properties) into "parent.child" keys one level deep.
func addNestedFieldTypes(out map[string]FieldTypeInfo, parent string, v *openapi3.Schema) {
	if v.Type == nil {
		return
	}
	if v.Type.Is("object") {
		addChildFieldTypes(out, parent, v.Properties)
		return
	}
	if v.Type.Is("array") && v.Items != nil && v.Items.Value != nil {
		addChildFieldTypes(out, parent, v.Items.Value.Properties)
	}
}
