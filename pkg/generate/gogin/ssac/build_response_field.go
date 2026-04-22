//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildResponseField — 응답 스키마의 단일 SchemaRef를 responseField로 변환

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// buildResponseField converts one OpenAPI SchemaRef into responseField
// metadata. Captures whether the property is a direct $ref or an array of
// $ref items so buildResponse can select the correct convert<Type>[List]
// shape when rendering a typed 200 response.
func buildResponseField(propName string, propRef *openapi3.SchemaRef, isRequired bool) responseField {
	rf := responseField{
		JSONName:   propName,
		GoName:     pascalCase(propName),
		IsRequired: isRequired,
	}
	if propRef == nil {
		return rf
	}
	if propRef.Ref != "" {
		rf.RefType = refName(propRef.Ref)
		return rf
	}
	if propRef.Value == nil {
		return rf
	}
	if !propRef.Value.Type.Is("array") {
		return rf
	}
	if propRef.Value.Items == nil || propRef.Value.Items.Ref == "" {
		return rf
	}
	rf.RefType = refName(propRef.Value.Items.Ref)
	rf.IsArray = true
	return rf
}
