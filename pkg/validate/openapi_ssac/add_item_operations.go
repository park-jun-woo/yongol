//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what addItemOperations — PathItem의 모든 Operation을 맵에 등록

package openapi_ssac

import "github.com/getkin/kin-openapi/openapi3"

// addItemOperations adds all operations of a path item to opMap by operationId.
func addItemOperations(opMap map[string]*openapi3.Operation, item *openapi3.PathItem) {
	for _, op := range []*openapi3.Operation{item.Get, item.Post, item.Put, item.Delete, item.Patch} {
		if op == nil || op.OperationID == "" {
			continue
		}
		opMap[op.OperationID] = op
	}
}
