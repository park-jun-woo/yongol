//ff:func feature=generate type=util control=iteration dimension=1
//ff:what operationId로 OpenAPI operation을 찾아 반환한다

package generate

import "github.com/getkin/kin-openapi/openapi3"

// findOpenAPIOpByID searches the OpenAPI document for an operation matching the
// given operationId. Returns the operation and true when found.
func findOpenAPIOpByID(doc *openapi3.T, operationID string) (*openapi3.Operation, bool) {
	if doc == nil || doc.Paths == nil {
		return nil, false
	}
	for _, item := range doc.Paths.Map() {
		if op, found := matchOperationByID(item, operationID); found {
			return op, true
		}
	}
	return nil, false
}
