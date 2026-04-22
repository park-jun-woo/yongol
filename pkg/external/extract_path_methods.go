//ff:func feature=external type=generator control=iteration dimension=1
//ff:what 단일 경로의 HTTP 메서드별 오퍼레이션을 추출한다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func extractPathMethods(pi *openapi3.PathItem, path string) []methodInfo {
	var methods []methodInfo
	for _, httpMethod := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
		op := pi.GetOperation(httpMethod)
		if op == nil || op.OperationID == "" {
			continue
		}
		methods = append(methods, buildMethodInfo(op, httpMethod, path))
	}
	return methods
}
