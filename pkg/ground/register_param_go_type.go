//ff:func feature=rule type=loader control=sequence topic=openapi
//ff:what registerParamGoType — OpenAPI 파라미터의 schema type+format 에서 Go 타입을 해석하여 Ground 에 등록

package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// registerParamGoType resolves the Go type for an OpenAPI parameter and
// registers it in g.Types under the "OpenAPI.paramType.<opID>.<name>" key.
func registerParamGoType(g *rule.Ground, opID string, p *openapi3.Parameter) {
	if p.Schema == nil || p.Schema.Value == nil {
		return
	}
	sv := p.Schema.Value
	if sv.Type == nil || len(*sv.Type) == 0 {
		return
	}
	goType := resolveOAPIParamGoType((*sv.Type)[0], sv.Format)
	if goType != "" {
		g.Types["OpenAPI.paramType."+opID+"."+p.Name] = goType
	}
}
