//ff:func feature=gen-hurl type=util control=sequence
//ff:what needsAuth — operation.Security 우선, nil이면 doc.Security 상속으로 인증 판정
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// needsAuth returns true when the operation requires an Authorization header.
//
// OpenAPI 3 inheritance rules:
//   - op.Security != nil (explicit override, even empty slice) wins.
//     A non-nil empty slice means "opt out of global security".
//   - op.Security == nil falls back to the document-level doc.Security.
//   - Everything nil / empty → false.
func needsAuth(op *openapi3.Operation, doc *openapi3.T) bool {
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
