//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what 단일 PathItem의 오퍼레이션들을 엔드포인트 목록에 추가한다

package react

import "github.com/getkin/kin-openapi/openapi3"

// appendOperations appends endpoints from a single path item.
func appendOperations(endpoints []endpoint, path string, pi *openapi3.PathItem) []endpoint {
	for method, op := range pi.Operations() {
		if op == nil || op.OperationID == "" {
			continue
		}
		endpoints = append(endpoints, endpoint{
			method:     method,
			path:       path,
			opID:       op.OperationID,
			pathParams: extractPathParams(path),
		})
	}
	return endpoints
}
