//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=config-check
//ff:what 테스트 헬퍼 — operationId→2xx 응답 필드명 매핑으로 최소 OpenAPI 문서 생성

package openapi_manifest

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildDocWithResponseFields builds a minimal OpenAPI doc whose ops each
// declare a 200 application/json object response with the given top-level
// property names. Used by the XON-60 unit tests.
func buildDocWithResponseFields(ops map[string][]string) *openapi3.T {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info:    &openapi3.Info{Title: "t", Version: "1"},
		Paths:   openapi3.NewPaths(),
	}
	for id, fields := range ops {
		pi := &openapi3.PathItem{
			Post: &openapi3.Operation{OperationID: id, Responses: build2xxObjectResponses(fields)},
		}
		doc.Paths.Set("/op-"+id, pi)
	}
	return doc
}
