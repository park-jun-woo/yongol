//ff:func feature=validate type=util control=sequence dimension=1 topic=response-body-required
//ff:what buildO05Diagnostic — O-5 진단 메시지·라인 구성

package openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// buildO05Diagnostic composes the [O-5] diagnostic for a single missing
// 4xx/5xx response body. The line falls back to the operation declaration
// because LineIndex does not currently expose per-response status lines.
func buildO05Diagnostic(status, opID, path string, lines *openapi.LineIndex) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    "api/openapi.yaml",
		Line:    lines.OperationLine(opID),
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[O-5] response " + status + " in operation " + opIDOrPath(opID, path) + " must declare content: application/json with schema",
		Advice:  "yongol does not emit empty error responses. Every 4xx/5xx response must carry a structured body (e.g. {error: string, code: string} schema). Add content: application/json + schema to the response definition. RFC 7807 Problem Details recommended.",
	}
}
