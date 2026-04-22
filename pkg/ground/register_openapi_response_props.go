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
	}
}
