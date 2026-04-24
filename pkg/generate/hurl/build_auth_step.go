//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildAuthStep — auth operation에서 step 구조체 조립 (token capture + isSignupRole 이면 smoke_email seed)
package hurl

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// buildAuthStep assembles a step from an auth operation.
// The captured tokenVar is recorded in ctx.captures for downstream resolution.
//
// BUG-015 / BUG-023 — Signup and Login bodies must reference the SAME
// email value. The default dummy generator inserts {{newUuid}} inline
// which hurl re-evaluates on every request, so Signup would land
// email=X while Login would send email=Y and 401. To keep the pair
// consistent, the signup-role step seeds a `smoke_email` variable via
// `[Options]` and the body is rewritten to reference that variable.
// Login inherits the same value because hurl variables set by
// `[Options]` persist across subsequent requests.
//
// The `isSignupRole` parameter replaced the previous hardcoded
// `EqualFold("Register")` check — Phase003 makes the pair role-driven
// so `Signup`, `Join`, `EnrollStudent`, etc. all seed the email
// correctly.
func buildAuthStep(method, path string, op *openapi3.Operation, ctx *scenarioCtx, role, tokenVar, sectionComment string, isFirst, isSignupRole bool) *step {
	body := generateRequestBody(op, ctx.fs, role)
	body = strings.ReplaceAll(body, "smoke-{{newUuid}}@example.com", "{{smoke_email}}")
	s := step{
		Method:      method,
		Path:        substitutePathParams(path, nil),
		OperationID: op.OperationID,
		NeedsAuth:   false,
		HasBody:     body != "",
		BodyJSON:    body,
		StatusCode:  inferSuccessStatus(op),
		Captures:    []capture{{VarName: tokenVar, JSONPath: `$.access_token`}},
	}
	if isSignupRole {
		s.Options = append(s.Options, "smoke_email=smoke-{{newUuid}}@example.com")
	}
	if sectionComment != "" && isFirst {
		s.Comment = sectionComment
	}
	ctx.captures[tokenVar] = true
	return &s
}
