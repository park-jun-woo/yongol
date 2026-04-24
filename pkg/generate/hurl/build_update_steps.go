//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what buildUpdateSteps — Phase3d PUT/PATCH endpoint step 생성
package hurl

import (
	"strings"
)

// buildUpdateSteps produces PUT/PATCH steps for CRUD resources.
func buildUpdateSteps(ctx *scenarioCtx) []step {
	var steps []step
	for path, pathItem := range ctx.fs.OpenAPIDoc.Paths.Map() {
		for method, op := range pathItem.Operations() {
			upper := strings.ToUpper(method)
			if upper != "PUT" && upper != "PATCH" {
				continue
			}
			if op.OperationID == "" || isAuthOpID(ctx, op.OperationID) {
				continue
			}
			if !canResolvePathParams(path, ctx.captures) {
				continue
			}
			body := generateRequestBody(op, ctx.fs, "")
			s := step{
				Method:      upper,
				Path:        substitutePathParams(path, ctx.captures),
				OperationID: op.OperationID,
				NeedsAuth:   needsAuth(op, ctx.fs.OpenAPIDoc),
				TokenVar:    resolveTokenVar(op.OperationID, ctx.roleMap, ctx.captures),
				HasBody:     body != "",
				BodyJSON:    body,
				StatusCode:  inferSuccessStatus(op),
				Assertions:  generateResponseAssertions(op),
			}
			steps = append(steps, s)
		}
	}
	return steps
}
