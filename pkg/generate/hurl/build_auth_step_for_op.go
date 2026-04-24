//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildAuthStepForOp — detectedAuthOp → *openapi3.Operation 해결 후 auth step 빌드

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildAuthStepForOp fetches the *openapi3.Operation for the detected
// auth op and delegates to buildAuthStep. Returns nil if the OpenAPI
// doc no longer contains the op (should not happen — detection uses the
// same doc — but kept defensive).
func buildAuthStepForOp(d *detectedAuthOp, ctx *scenarioCtx, role, tokenVar, sectionComment string, isFirst, isSignupRole bool) *step {
	pathItem := ctx.fs.OpenAPIDoc.Paths.Find(d.Path)
	if pathItem == nil {
		return nil
	}
	var op *openapi3.Operation
	for method, o := range pathItem.Operations() {
		if method == d.Method && o != nil && o.OperationID == d.OpID {
			op = o
			break
		}
	}
	if op == nil {
		return nil
	}
	return buildAuthStep(d.Method, d.Path, op, ctx, role, tokenVar, sectionComment, isFirst, isSignupRole)
}
