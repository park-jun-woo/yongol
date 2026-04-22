//ff:func feature=gen-gogin type=util control=sequence
//ff:what requiresAuth — operation.Security 우선, nil이면 doc.Security 상속 판정
package boot

import "github.com/getkin/kin-openapi/openapi3"

// requiresAuth applies OpenAPI 3 security inheritance rules:
//   - op.Security != nil → use it as-is (len 0 means explicit opt-out)
//   - op.Security == nil → inherit from doc.Security
//
// Mirrors pkg/generate/hurl/needs_auth.go (needsAuth) — kept in sync manually
// until a shared security helper is extracted.
func requiresAuth(op *openapi3.Operation, doc *openapi3.T) bool {
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
