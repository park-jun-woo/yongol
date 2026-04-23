//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what findAuthOp — OpenAPI 전체에서 지정 operationID(Register/Login) 단일 매칭

package hurl

import (
	"strings"
)

// findAuthOp walks the parsed OpenAPI paths looking for a single
// operationId that equals wantOpID (case-insensitive). On match, it
// delegates to buildAuthStep to assemble the hurl step. Returns nil
// when the wanted auth op is not declared in OpenAPI — callers treat
// that as "skip this auth step".
//
// findAuthOp is invoked once per expected opID from authOpOrder, so
// the emitted auth steps appear in a deterministic order regardless of
// how the underlying OpenAPI map iterates.
func findAuthOp(ctx *scenarioCtx, wantOpID, role, tokenVar, sectionComment string, isFirst bool) *step {
	for path, pathItem := range ctx.fs.OpenAPIDoc.Paths.Map() {
		for method, op := range pathItem.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			if !strings.EqualFold(op.OperationID, wantOpID) {
				continue
			}
			return buildAuthStep(method, path, op, ctx, role, tokenVar, sectionComment, isFirst)
		}
	}
	return nil
}
