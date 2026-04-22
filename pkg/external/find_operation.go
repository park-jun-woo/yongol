//ff:func feature=external type=util control=sequence
//ff:what OpenAPI 문서에서 특정 HTTP 메서드와 경로의 오퍼레이션을 찾는다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func findOperation(doc *openapi3.T, method, path string) *openapi3.Operation {
	if doc.Paths == nil {
		return nil
	}
	pi := doc.Paths.Find(path)
	if pi == nil {
		return nil
	}
	return pi.GetOperation(method)
}
