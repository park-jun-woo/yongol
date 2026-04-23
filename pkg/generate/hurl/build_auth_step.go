//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildAuthStep — auth operation에서 step 구조체 조립 (token capture + smoke_email 공유)
package hurl

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// buildAuthStep assembles a step from an auth operation.
// The captured tokenVar is recorded in ctx.captures for downstream resolution.
//
// BUG-015 / Phase003 — Register and Login bodies must reference the
// SAME email value. The default dummy generator inserts {{newUuid}}
// inline which hurl re-evaluates on every request, so Register would
// land email=X while Login would send email=Y and 401. To keep the
// pair consistent, the Register step seeds a `smoke_email` variable
// via `[Options]` and the body is rewritten to reference that
// variable. Login inherits the same value because hurl variables set
// by `[Options]` persist across subsequent requests.
func buildAuthStep(method, path string, op *openapi3.Operation, ctx *scenarioCtx, role, tokenVar, sectionComment string, isFirst bool) *step {
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
	if strings.EqualFold(op.OperationID, "Register") {
		s.Options = append(s.Options, "smoke_email=smoke-{{newUuid}}@example.com")
	}
	if sectionComment != "" && isFirst {
		s.Comment = sectionComment
	}
	ctx.captures[tokenVar] = true
	return &s
}
