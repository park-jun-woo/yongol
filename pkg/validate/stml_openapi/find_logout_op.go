//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-openapi
//ff:what FindLogoutOp — OpenAPI doc에서 logout-like operationId 검색 (security 선언 무관)

package stml_openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// FindLogoutOp returns the operationId of the first logout-like operation
// (case-insensitive "logout" in operationId), or "" if none. The search
// ignores the operation's own security declaration: bearer/JWT logout
// endpoints commonly use security: [] because the access token may be
// expired at logout time (the refresh token is sent in the body instead).
// Callers (resolveLayoutLogoutOps, tm58BearerLogoutOpHint) already verify
// that the project uses bearer auth, so the operation-level auth check is
// unnecessary and was incorrectly filtering out security: [] logout ops.
func FindLogoutOp(doc *openapi3.T) string {
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
		if strings.Contains(strings.ToLower(op.OperationID), "logout") {
			return op.OperationID
		}
	}
	return ""
}
