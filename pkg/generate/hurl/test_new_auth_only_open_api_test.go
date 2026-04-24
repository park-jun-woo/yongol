//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newAuthOnlyOpenAPI — Register+Login 두 operation 만 포함한 테스트용 *openapi3.T 생성 (security: [] + password field)

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// newAuthOnlyOpenAPI returns an *openapi3.T with just Register + Login
// auth operations wired under /auth/register and /auth/login. Both ops
// declare `security: []` (explicit public override) and carry a
// `password` request-body field so detectAuthOps (Phase003 shape
// detection) can classify them once the paired SSaC ServiceFuncs are
// attached via newSmokeFullstack.
//
// Callers typically attach this to Fullstack.OpenAPIDoc and run the
// smoke order builder to assert auth-step ordering.
func newAuthOnlyOpenAPI() *openapi3.T {
	return newAuthOpenAPI("Register", "Login")
}
