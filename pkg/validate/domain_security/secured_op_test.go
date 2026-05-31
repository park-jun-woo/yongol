//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestByName_ZeroCov — domain_security 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package domain_security

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func securedOp(opID string) *openapi3.Operation {
	sr := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearer": {}}}
	return &openapi3.Operation{OperationID: opID, Security: &sr}
}
