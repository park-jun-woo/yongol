//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-openapi
//ff:what findLogoutOp — OpenAPI doc에서 auth 필수인 logout-like operationId 검색

package stml_openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// findLogoutOp returns the operationId of the first logout-like operation
// (case-insensitive "logout" in operationId) that requires auth, or "" if none.
func findLogoutOp(doc *openapi3.T) string {
	if doc.Paths == nil {
		return ""
	}
	var ops []*openapi3.Operation
	for _, item := range doc.Paths.Map() {
		ops = append(ops, item.Post, item.Put, item.Delete, item.Patch)
	}
	for _, op := range ops {
		if op == nil || op.OperationID == "" {
			continue
		}
		if strings.Contains(strings.ToLower(op.OperationID), "logout") &&
			OpRequiresAuth(op, doc) {
			return op.OperationID
		}
	}
	return ""
}
