//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XSO-16 — verifies that every OpenAPI operationId is implemented as an SSaC funcName

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
			Advice:  "Add an SSaC function corresponding to operationId " + opID,
		})
	}
	return diags
}
