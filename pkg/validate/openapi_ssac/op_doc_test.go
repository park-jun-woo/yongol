//ff:func feature=validate type=test-helper control=sequence topic=ssac-openapi
//ff:what opDoc — operationId 하나를 가진 단일 경로 POST OpenAPI doc 생성

package openapi_ssac

import "github.com/getkin/kin-openapi/openapi3"

// opDoc builds a minimal OpenAPI doc with one POST operation under path.
func opDoc(path, opID string) *openapi3.T {
	return &openapi3.T{
		Paths: openapi3.NewPaths(openapi3.WithPath(path, &openapi3.PathItem{
			Post: &openapi3.Operation{OperationID: opID},
		})),
	}
}
