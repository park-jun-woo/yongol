//ff:func feature=validate type=util control=iteration dimension=1 topic=tsx-openapi
//ff:what XOT-2 헬퍼 — 단일 apiClient 호출의 인자 키를 OpenAPI parameters 와 대조

package tsx_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	tsxparser "github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// xot02CheckCall validates every argument key of a single apiClient call
// against the OpenAPI parameter set for the called operationId.
func xot02CheckCall(file string, call tsxparser.APICall, params rule.StringSet) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, arg := range call.Args {
		if isTransportKey(arg.Key) {
			continue
		}
		if params[arg.Key] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Line:    call.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XOT-2] apiClient." + call.OperationID + "({" + arg.Key + ": ...}) — '" + arg.Key + "' is not declared as an OpenAPI parameter",
			Advice:  "Add " + arg.Key + " to the parameters of the operation in openapi.yaml, or correct the argument name in the call",
		})
	}
	return diags
}
