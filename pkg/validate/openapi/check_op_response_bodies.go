//ff:func feature=validate type=util control=iteration dimension=1 topic=response-body-required
//ff:what checkOpResponseBodies — 단일 operation 의 4xx/5xx response body 유무 검사

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// checkOpResponseBodies emits one [O-5] diagnostic per 4xx/5xx response in op
// that is missing `content: application/json + schema`. Helper extracted from
// o05ResponseBodyRequired to keep loop nesting under filefunc Q1's 3-level
// limit.
func checkOpResponseBodies(op *openapi3.Operation, path string, lines *openapi.LineIndex) []diagnostic.Diagnostic {
	if op == nil || op.Responses == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for status, resp := range op.Responses.Map() {
		if !isErrorStatus(status) {
			continue
		}
		if hasJSONContentWithSchema(resp) {
			continue
		}
		diags = append(diags, buildO05Diagnostic(status, op.OperationID, path, lines))
	}
	return diags
}
