//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what registerOpenAPIResponseProps — 단일 2xx schema properties 를 Ground.Types 에 등록
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// registerOpenAPIResponseProps walks a single schema's top-level properties
// and registers each field's Go type under "OpenAPI.response.<opID>.<field>".
func registerOpenAPIResponseProps(g *rule.Ground, opID string, schema *openapi3.Schema) {
	if schema == nil {
		return
	}
	for propName, propRef := range schema.Properties {
		t := responsePropType(propRef)
		if t == "" {
			continue
		}
		g.Types["OpenAPI.response."+opID+"."+propName] = t
	}
}

//ff:func feature=rule type=util control=selection
//ff:what responsePropType — 응답 본문 필드 1개의 oapi-codegen 생성 Go 타입 결정 (format-aware)

// responsePropType resolves a single response-body property to the Go type
// oapi-codegen generates. Primitive `string` schemas (no $ref) are mapped
// format-aware via resolveOAPIResponseGoType (uuid→openapi_types.UUID,
// date-time→time.Time, email→openapi_types.Email, else string), so XOS-67
// compares @response values against the type the generated struct field
// actually has. $ref / integer / array / object schemas keep the existing
// resolveSchemaType path.
func responsePropType(ref *openapi3.SchemaRef) string {
	if ref != nil && ref.Ref == "" && ref.Value != nil {
		if types := ref.Value.Type; types != nil && len(types.Slice()) > 0 && types.Slice()[0] == "string" {
			return resolveOAPIResponseGoType("string", ref.Value.Format)
		}
	}
	return resolveSchemaType(ref)
}
