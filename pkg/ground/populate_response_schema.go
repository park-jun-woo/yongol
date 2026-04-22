//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateResponseSchema — operationId별 OpenAPI response 필드를 Ground.Schemas에 등록
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

		// 상태 코드별 키: OpenAPI.response.<code>.<opID>
		g.Schemas["OpenAPI.response."+code+"."+opID] = fields

		// $ref 解決 (상태 코드별)
		resolved := resolveRefProperties(ct.Schema.Value)
		if len(resolved) > 0 {
			g.Schemas["OpenAPI.response.resolved."+code+"."+opID] = resolved
		}

		// 첫 2xx 는 기존 키(OpenAPI.response.<opID>)로도 유지 — 하위 호환
		if !primary2xxDone && code[0] == '2' {
			g.Schemas["OpenAPI.response."+opID] = fields
			if len(resolved) > 0 {
				g.Schemas["OpenAPI.response.resolved."+opID] = resolved
			}
			primary2xxDone = true
		}
	}
}
