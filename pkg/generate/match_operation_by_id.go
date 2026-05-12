//ff:func feature=generate type=util control=iteration dimension=1
//ff:what PathItem의 operations에서 operationId가 일치하는 operation을 찾는다
package generate

import "github.com/getkin/kin-openapi/openapi3"

// matchOperationByID iterates the operations of a PathItem and returns the one
// whose OperationID matches the given id.
func matchOperationByID(item *openapi3.PathItem, operationID string) (*openapi3.Operation, bool) {
	for _, op := range item.Operations() {
		if op.OperationID == operationID {
			return op, true
		}
	}
	return nil, false
}
