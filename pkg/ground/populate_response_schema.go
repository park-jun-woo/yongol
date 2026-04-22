//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateResponseSchema — registers OpenAPI response fields per operationId into Ground.Schemas
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateResponseSchema(g *rule.Ground, opID string, op *openapi3.Operation) {
	if op.Responses == nil {
		return
	}
	primary2xxDone := false
	for code, resp := range op.Responses.Map() {
		if len(code) == 0 {
			continue
		}
		switch code[0] {
		case '2', '4', '5':
		default:
			continue
		}
		if resp.Value == nil || resp.Value.Content == nil {
			continue
		}
		ct := resp.Value.Content.Get("application/json")
		if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
			continue
		}
		var fields []string
		for name := range ct.Schema.Value.Properties {
			fields = append(fields, name)
		}

		// per-status-code key: OpenAPI.response.<code>.<opID>
		g.Schemas["OpenAPI.response."+code+"."+opID] = fields

		// $ref resolution (per status code)
		resolved := resolveRefProperties(ct.Schema.Value)
		if len(resolved) > 0 {
			g.Schemas["OpenAPI.response.resolved."+code+"."+opID] = resolved
		}

		// also keep the first 2xx under the legacy key (OpenAPI.response.<opID>) for back-compat
		if !primary2xxDone && code[0] == '2' {
			g.Schemas["OpenAPI.response."+opID] = fields
			if len(resolved) > 0 {
				g.Schemas["OpenAPI.response.resolved."+opID] = resolved
			}
			primary2xxDone = true
		}
	}
}
