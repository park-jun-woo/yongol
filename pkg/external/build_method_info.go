//ff:func feature=external type=generator control=sequence
//ff:what 단일 오퍼레이션으로부터 메서드 정보를 구성한다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func buildMethodInfo(op *openapi3.Operation, httpMethod, path string) methodInfo {
	m := methodInfo{
		Name:       toPascalCase(op.OperationID),
		HTTPMethod: httpMethod,
		Path:       path,
	}

	m.Params = append(m.Params, extractPathParams(op)...)
	m.Params = append(m.Params, extractBodyParams(op)...)
	m.ReturnType = detectReturnType(op)

	return m
}
