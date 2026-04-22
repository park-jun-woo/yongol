//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildAuthStep — auth operation에서 step 구조체 조립 (token capture 포함)
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildAuthStep assembles a step from an auth operation.
// The captured tokenVar is recorded in ctx.captures for downstream resolution.
func buildAuthStep(method, path string, op *openapi3.Operation, ctx *scenarioCtx, role, tokenVar, sectionComment string, isFirst bool) *step {
	body := generateRequestBody(op, ctx.fs, role)
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
	if sectionComment != "" && isFirst {
		s.Comment = sectionComment
	}
	ctx.captures[tokenVar] = true
	return &s
}
