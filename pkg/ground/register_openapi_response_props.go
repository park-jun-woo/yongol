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
		t := resolveSchemaType(propRef)
		if t == "" {
			continue
		}
		g.Types["OpenAPI.response."+opID+"."+propName] = t
		// Preserve the date-time format as a sibling marker so XOS-67 can
		// recognise a DDL TIMESTAMPTZ (time.Time) bound to an OpenAPI
		// { type: string, format: date-time } field as compatible. The
		// primary type value stays "string" (resolvePrimitiveType drops
		// string formats by design) to keep literal-string bindings valid.
		if t == "string" && propRef != nil && propRef.Value != nil && propRef.Value.Format == "date-time" {
			g.Types["OpenAPI.response."+opID+"."+propName+".format"] = "date-time"
		}
	}
}
