//ff:func feature=validate type=util control=sequence topic=domain-security
//ff:what hasNonEmptySecurity — operation에 비어있지 않은 security 선언이 있는지 확인
package domain_security

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// hasNonEmptySecurity returns true when the operation has an explicit non-empty
// security requirement (at least one entry).
func hasNonEmptySecurity(op *openapi3.Operation) bool {
	return op.Security != nil && len(*op.Security) > 0
}
