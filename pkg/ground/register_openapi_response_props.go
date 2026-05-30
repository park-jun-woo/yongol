//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what registerOpenAPIResponseProps — 단일 2xx schema properties 를 Ground.Types 에 등록
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// registerOpenAPIResponseProps walks a single schema's top-level properties
// and registers each field's Go type under "OpenAPI.response.<opID>.<field>".
// Each property is resolved by the unified resolveOAPIGoType in the
// CtxResponseBody context, so string/array formats (uuid→openapi_types.UUID,
// date-time→time.Time, []uuid→[]openapi_types.UUID, …) are honoured the same
// way oapi-codegen renders the generated struct field. XOS-67 consumes these
// values to compare @response literals against the field's actual Go type.
func registerOpenAPIResponseProps(g *rule.Ground, opID string, schema *openapi3.Schema) {
	if schema == nil {
		return
	}
	for propName, propRef := range schema.Properties {
		t := resolveOAPIGoType(propRef, CtxResponseBody)
		if t == "" {
			continue
		}
		g.Types["OpenAPI.response."+opID+"."+propName] = t
	}
}
