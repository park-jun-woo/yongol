//ff:func feature=validate type=test-helper control=sequence topic=response-body-required
//ff:what 테스트 헬퍼 — description 만 있는 (content 없는) ResponseRef 빌드

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// emptyResponse builds a ResponseRef with description only — no content.
// Used by O-5 unit tests to simulate the BUG-040 reproduce condition.
func emptyResponse(desc string) *openapi3.ResponseRef {
	d := desc
	return &openapi3.ResponseRef{
		Value: &openapi3.Response{Description: &d},
	}
}
