//ff:func feature=gen-ir type=util control=sequence
//ff:what paramSchemaType -- OpenAPI 파라미터의 첫 번째 schema type 문자열 추출 (없으면 "")

package ir

import "github.com/getkin/kin-openapi/openapi3"

// paramSchemaType returns the first schema type string of an OpenAPI parameter,
// or "" when the schema or its type is absent.
func paramSchemaType(p *openapi3.ParameterRef) string {
	if p.Value.Schema == nil || p.Value.Schema.Value == nil || p.Value.Schema.Value.Type == nil {
		return ""
	}
	return p.Value.Schema.Value.Type.Slice()[0]
}
