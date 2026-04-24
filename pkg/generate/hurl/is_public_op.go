//ff:func feature=gen-hurl type=util control=sequence
//ff:what isPublicOp — OpenAPI op 이 security: [] 로 공개 선언되었는지 판정 (nil 은 global 상속)

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// isPublicOp returns true when the operation explicitly opts out of
// global security via `security: []` (a non-nil, zero-length slice).
// A nil op.Security means "inherit doc.Security" — NOT public — so
// those ops are excluded from auth-shape detection.
//
// The OpenAPI 3 spec uses a non-nil empty array to override global
// security to "no auth required"; this is the standard signal used by
// auth endpoints (signup / login) that must be callable without a
// token.
func isPublicOp(op *openapi3.Operation) bool {
	if op == nil {
		return false
	}
	return op.Security != nil && len(*op.Security) == 0
}
