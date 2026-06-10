//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what OpRequiresAuth — operation.Security 우선, nil이면 doc.Security 상속으로 보호 여부 판정

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// OpRequiresAuth applies OpenAPI 3 security inheritance rules:
//   - op.Security != nil → use it as-is (len 0 means explicit opt-out)
//   - op.Security == nil → inherit from doc.Security
//
// Mirrors pkg/generate/gogin/boot/requires_auth.go (requiresAuth) — kept
// in sync manually until a shared security helper is extracted. Exported
// (Phase005) so pkg/generate/react derives protected routes from the same
// security judgment the XMO/TM rules use.
func OpRequiresAuth(op *openapi3.Operation, doc *openapi3.T) bool {
	if op == nil {
		return false
	}
	if op.Security != nil {
		return len(*op.Security) > 0
	}
	if doc == nil {
		return false
	}
	return len(doc.Security) > 0
}
