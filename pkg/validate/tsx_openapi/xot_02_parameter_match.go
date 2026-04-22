//ff:func feature=validate type=rule control=iteration dimension=2 topic=tsx-openapi
//ff:what XOT-2 — apiClient 호출의 path/query 파라미터 객체 키가 OpenAPI parameters 에 존재하는지 검증
package tsx_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xot02ParameterMatch validates XOT-2: for every apiClient.<op>({...})
// invocation, each top-level key of the argument object must match a path
// or query parameter declared on the OpenAPI operation.
//
// Keys that yongol / the generated api.ts uses as transport containers
// (e.g. `body`, `data`) are intentionally ignored so mutation-style calls
// (createWorkflow({ body: {...} })) don't produce false positives. Request
// body fields are the domain of XOT-3, not this rule.
//
// Skipped when XOT-1 would already fire for the same operation — since a
// missing operationId has no parameter set to compare against.
func xot02ParameterMatch(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.TSXPages) == 0 {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	opIDs := g.Lookup["OpenAPI.operationId"]
	var diags []diagnostic.Diagnostic
	for _, page := range fs.TSXPages {
		for _, call := range page.Calls {
			if !opIDs[call.OperationID] {
				continue // XOT-1 already reports this
			}
			params := g.Lookup["OpenAPI.param."+call.OperationID]
			for _, arg := range call.Args {
				if isTransportKey(arg.Key) {
					continue
				}
				if params[arg.Key] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    page.File,
					Line:    call.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[XOT-2] apiClient." + call.OperationID + "({" + arg.Key + ": ...}) — '" + arg.Key + "' is not declared as an OpenAPI parameter",
					Advice:  "openapi.yaml 에서 해당 operation 의 parameters 에 " + arg.Key + " 를 추가하거나 호출 인자명을 수정하세요",
				})
			}
		}
	}
	return diags
}

// isTransportKey returns true for argument keys that are wrappers around
// the request body rather than actual OpenAPI parameters. XOT-3 validates
// the body contents; XOT-2 must skip them to avoid double-counting.
func isTransportKey(key string) bool {
	switch key {
	case "body", "data", "payload", "json":
		return true
	}
	return false
}
