//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XSO-16 — OpenAPI operationId가 SSaC funcName에 구현되어 있는지 검증

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xso16OpIDToFunc validates XSO-16: every OpenAPI operationId has a matching
// SSaC funcName.
func xso16OpIDToFunc(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	funcs := g.Lookup["SSaC.funcName"]
	var diags []diagnostic.Diagnostic
	for opID := range g.Lookup["OpenAPI.operationId"] {
		if funcs[opID] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "api/openapi.yaml",
			Line:    fs.OpenAPILines.OperationLine(opID),
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XSO-16] OpenAPI operationId " + opID + " has no matching SSaC func",
			Advice:  "operationId " + opID + " 에 대응하는 SSaC 함수를 추가하세요",
		})
	}
	return diags
}
