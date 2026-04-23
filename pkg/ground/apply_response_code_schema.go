//ff:func feature=rule type=loader control=sequence
//ff:what applyResponseCodeSchema — 단일 status code 의 JSON response 를 Ground 에 반영
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// applyResponseCodeSchema writes the per-status-code and (for the first 2xx)
// back-compat schema keys for a single OpenAPI response. Returns the updated
// primary2xxDone flag so the caller loop can short-circuit subsequent 2xx
// primary writes.
func applyResponseCodeSchema(g *rule.Ground, opID, code string, resp *openapi3.ResponseRef, primary2xxDone bool) bool {
	if len(code) == 0 {
		return primary2xxDone
	}
	if !isResponseCodeRelevant(code) {
		return primary2xxDone
	}
	if resp.Value == nil || resp.Value.Content == nil {
		return primary2xxDone
	}
	ct := resp.Value.Content.Get("application/json")
	if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
		return primary2xxDone
	}
	fields := collectSchemaFieldNames(ct.Schema.Value)

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
		return true
	}
	return primary2xxDone
}
