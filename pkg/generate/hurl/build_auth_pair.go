//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildAuthPair — detectAuthOps 결과로 signup→login 순서 step 생성 (BUG-015/023: 빈 DB 실행 가능)

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildAuthPair builds auth steps in the order signup → login using
// the shape-detected auth ops stored on ctx (populated by detectAuthOps
// in newScenarioCtx). Signup always comes first so:
//
//   - Empty DB → signup 201 creates a fresh account; login then verifies.
//   - Both responses capture `access_token` into the same variable; the
//     login-issued token overrides signup's, so downstream CRUD steps
//     always run with the freshest credential.
//
// Returns nil when OpenAPIDoc is missing; emits only whichever side was
// detected (both/one/neither). Pre-Phase003 the ordering was driven by
// a hardcoded `{"Register", "Login"}` name list which silently broke
// any project using a different operationId (Signup, Join, SignIn, …).
func buildAuthPair(ctx *scenarioCtx, role, tokenVar, sectionComment string) []step {
	if ctx.fs.OpenAPIDoc == nil {
		return nil
	}
	var out []step
	if ctx.authSignup != nil {
		if s := buildAuthStepForOp(ctx.authSignup, ctx, role, tokenVar, sectionComment, len(out) == 0, true); s != nil {
			out = append(out, *s)
		}
	}
	if ctx.authLogin != nil {
		if s := buildAuthStepForOp(ctx.authLogin, ctx, role, tokenVar, sectionComment, len(out) == 0, false); s != nil {
			out = append(out, *s)
		}
	}
	return out
}

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
