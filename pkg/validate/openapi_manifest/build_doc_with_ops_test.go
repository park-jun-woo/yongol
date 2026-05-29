//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=config-check
//ff:what 테스트 헬퍼 — operationId 리스트로 최소 OpenAPI 문서 생성

package openapi_manifest

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func buildDocWithOps(opIDs []string) *openapi3.T {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info:    &openapi3.Info{Title: "t", Version: "1"},
		Paths:   openapi3.NewPaths(),
	}
	for i, id := range opIDs {
		path := "/ep" + string(rune('a'+i))
		pi := &openapi3.PathItem{
			Post: &openapi3.Operation{OperationID: id, Responses: openapi3.NewResponses()},
		}
		doc.Paths.Set(path, pi)
	}
	return doc
}
