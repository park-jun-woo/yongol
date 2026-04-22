//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what matchAuthOp — pathItem에서 Register/Login operation을 찾아 step 생성
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// matchAuthOp finds a Register or Login operation in the pathItem and returns a step.
func matchAuthOp(pathItem *openapi3.PathItem, path string, ctx *scenarioCtx, role, tokenVar, sectionComment string, isFirst bool) *step {
	for method, op := range pathItem.Operations() {
		if op.OperationID == "" || !isAuthOpID(op.OperationID) {
			continue
		}
		return buildAuthStep(method, path, op, ctx, role, tokenVar, sectionComment, isFirst)
	}
	return nil
}
