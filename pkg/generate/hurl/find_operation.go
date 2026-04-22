//ff:func feature=gen-hurl type=util control=sequence
//ff:what findOperation — OpenAPI doc에서 path+method로 Operation 조회
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// findOperation returns the Operation for a given path and method.
func findOperation(doc *openapi3.T, path, method string) *openapi3.Operation {
	pathItem := doc.Paths.Find(path)
	if pathItem == nil {
		return nil
	}
	return pathItem.Operations()[method]
}
