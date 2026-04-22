//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildPostStep — POST operation에서 step 구조체 조립 (TokenVar + capture 추가)
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildPostStep assembles a step from a POST operation.
// Resolves TokenVar from ctx.roleMap + ctx.captures, records any new capture.
func buildPostStep(op *openapi3.Operation, path, resource string, ctx *scenarioCtx) step {
	body := generateRequestBody(op, ctx.fs, "")
	captures := inferCaptureField(op, resource)
	s := step{
		Method:      "POST",
		Path:        substitutePathParams(path, ctx.captures),
		OperationID: op.OperationID,
		NeedsAuth:   needsAuth(op, ctx.fs.OpenAPIDoc),
		TokenVar:    resolveTokenVar(op.OperationID, ctx.roleMap, ctx.captures),
		HasBody:     body != "",
		BodyJSON:    body,
		StatusCode:  inferSuccessStatus(op),
		Captures:    captures,
		Assertions:  generateResponseAssertions(op),
	}
	recordCaptureVars(ctx, captures)
	return s
}
