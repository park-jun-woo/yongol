//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what resolveResponseEntity — operation 2xx 응답을 canonical 엔티티 component 로 해석 (SSaC 우선, 없으면 fallback)

package openapi_ddl

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// resolveResponseEntity returns the canonical entity component for an
// operation's 2xx response, or "" when the response is not a single-resource
// entity. When the operation has an SSaC @response sequence it is trusted fully
// (strategies A/B-1); otherwise the response schema itself is inspected
// (fallback). fn may be nil for operations without a SSaC function.
func resolveResponseEntity(idx *entityIndex, fn *ssac.ServiceFunc, schemaRef *openapi3.SchemaRef) string {
	if fn != nil {
		if seq := responseSeqOf(fn); seq != nil {
			return resolveSSaCEntity(idx, fn, seq)
		}
	}
	return resolveFallbackEntity(idx, schemaRef)
}
