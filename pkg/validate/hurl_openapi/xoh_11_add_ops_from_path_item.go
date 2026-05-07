//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what addOpsFromPathItem — PathItem 의 operation 들에서 operationId 를 ids 에 추가

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

func addOpsFromPathItem(ids map[string]bool, pi *openapi3.PathItem) {
	for _, op := range pi.Operations() {
		if op.OperationID != "" {
			ids[op.OperationID] = true
		}
	}
}
