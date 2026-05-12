//ff:func feature=validate type=util control=sequence topic=domain-security
//ff:what hasEmptySecurity — operation에 security: [] (빈 보안 설정)이 있는지 확인
package domain_security

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// hasEmptySecurity returns true when the operation has explicit `security: []`.
func hasEmptySecurity(op *openapi3.Operation) bool {
	return op.Security != nil && len(*op.Security) == 0
}
