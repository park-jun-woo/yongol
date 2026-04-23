//ff:func feature=openapi-parse type=test-helper control=iteration dimension=1
//ff:what opWith — 주어진 status 코드 목록으로 테스트용 Operation 생성

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// opWith builds a minimal *openapi3.Operation whose Responses contain exactly
// the status codes in codes. Used by DeriveSuccessStatus unit tests to
// construct compact fixtures.
func opWith(codes ...string) *openapi3.Operation {
	resp := openapi3.NewResponses()
	for _, c := range codes {
		resp.Set(c, &openapi3.ResponseRef{Value: openapi3.NewResponse()})
	}
	return &openapi3.Operation{Responses: resp}
}
