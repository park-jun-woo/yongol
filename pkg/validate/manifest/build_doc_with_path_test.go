//ff:func feature=validate type=test-helper control=sequence topic=manifest-observability
//ff:what buildDocWithPath — 테스트용 단일 path OpenAPI 문서 생성

package manifest

import "github.com/getkin/kin-openapi/openapi3"

// buildDocWithPath returns a minimal OpenAPI 3.0 document containing exactly
// one path (GET Any). Used by OBS-002 tests to exercise path-collision logic.
func buildDocWithPath(path string) *openapi3.T {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info:    &openapi3.Info{Title: "t", Version: "1"},
		Paths:   openapi3.NewPaths(),
	}
	pi := &openapi3.PathItem{
		Get: &openapi3.Operation{OperationID: "Any", Responses: openapi3.NewResponses()},
	}
	doc.Paths.Set(path, pi)
	return doc
}
