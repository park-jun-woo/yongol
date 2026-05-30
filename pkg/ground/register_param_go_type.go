//ff:func feature=rule type=loader control=sequence topic=openapi
//ff:what registerParamGoType — OpenAPI 파라미터의 schema type+format 에서 Go 타입을 해석하여 Ground 에 등록

package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// registerParamGoType resolves the Go type for an OpenAPI parameter and
// registers it in g.Types under the "OpenAPI.paramType.<opID>.<name>" key.
// The unified resolveOAPIGoType (CtxParam context) is used, so array
// parameters now register as []T (e.g. []openapi_types.UUID) instead of being
// skipped. xfs_73 is the sole type-value consumer of this key.
func registerParamGoType(g *rule.Ground, opID string, p *openapi3.Parameter) {
	goType := resolveOAPIGoType(p.Schema, CtxParam)
	if goType != "" {
		g.Types["OpenAPI.paramType."+opID+"."+p.Name] = goType
	}
}
