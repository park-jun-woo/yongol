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
// actually has.
//
// Arrays whose items are a primitive `string` are descended into so that
// { type: array, items: { type: string, format: uuid } } maps to
// []openapi_types.UUID (not []string). resolvePrimitiveType's array case
// drops the item format, which produced an XOS-67 false positive for
// array-of-uuid response fields (BUG-102). Only primitive-string items are
// format-aware here; $ref / nested-array / non-string items keep the
// resolveSchemaType path, as do $ref / integer / object top-level schemas.
func responsePropType(ref *openapi3.SchemaRef) string {
	if ref != nil && ref.Ref == "" && ref.Value != nil {
		if types := ref.Value.Type; types != nil && len(types.Slice()) > 0 {
			switch types.Slice()[0] {
			case "string":
				return resolveOAPIResponseGoType("string", ref.Value.Format)
			case "array":
				if items := ref.Value.Items; items != nil && items.Ref == "" && items.Value != nil {
					if it := items.Value.Type; it != nil && len(it.Slice()) > 0 && it.Slice()[0] == "string" {
						return "[]" + resolveOAPIResponseGoType("string", items.Value.Format)
					}
				}
			}
		}
	}
	return resolveSchemaType(ref)
}
