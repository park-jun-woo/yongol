//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what isAuthEndpoint — operation 이 security: [] 으로 인증 불필요 엔드포인트인지 판별

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// isAuthEndpoint returns true if the operation has an explicit empty
// security requirement (security: []), indicating it is an auth endpoint
// that does not require authentication.
func isAuthEndpoint(op *openapi3.Operation) bool {
	return op.Security != nil && len(*op.Security) == 0
}
