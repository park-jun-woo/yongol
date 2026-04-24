//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what addAuthFKFields — pathItem의 auth op body에서 _id 필드를 needed에 추가
package hurl

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// addAuthFKFields adds FK resource names from auth operations in the pathItem.
// ctx.authOpIDs (populated by detectAuthOps) decides which operations
// count as auth — no more Register/Login name literals.
func addAuthFKFields(ctx *scenarioCtx, pathItem *openapi3.PathItem, needed map[string]bool) {
	for _, op := range pathItem.Operations() {
		if op.OperationID == "" || !isAuthOpID(ctx, op.OperationID) {
			continue
		}
		for _, name := range extractBodyFieldNames(op) {
			if strings.HasSuffix(name, "_id") {
				needed[strings.TrimSuffix(name, "_id")] = true
			}
		}
	}
}
