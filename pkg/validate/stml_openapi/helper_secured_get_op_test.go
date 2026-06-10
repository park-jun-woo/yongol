//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what securedGetOp — security가 걸린 테스트용 GET operation PathItem 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// securedGetOp creates a PathItem with a GET operation that declares a
// non-empty security requirement (a protected endpoint).
func securedGetOp(opID string) *openapi3.PathItem {
	sec := openapi3.SecurityRequirements{openapi3.SecurityRequirement{"bearerAuth": {}}}
	return &openapi3.PathItem{Get: &openapi3.Operation{OperationID: opID, Security: &sec}}
}
