//ff:func feature=validate type=rule control=sequence topic=openapi-structural
//ff:what schemaPropIsString — 스키마에 OpenAPI string 타입의 지정 프로퍼티가 있는지 판정한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// schemaPropIsString reports whether schema's property field exists and is
// typed as an OpenAPI string.
func schemaPropIsString(schema *openapi3.Schema, field string) bool {
	prop, ok := schema.Properties[field]
	if !ok || prop == nil || prop.Value == nil || prop.Value.Type == nil {
		return false
	}
	return prop.Value.Type.Is("string")
}
