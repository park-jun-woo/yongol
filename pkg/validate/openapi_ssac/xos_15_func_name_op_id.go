//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-15 — verifies that every SSaC funcName is defined as an OpenAPI operationId

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos15FuncNameOpID validates XOS-15: every SSaC funcName (non-subscribe)
// has a matching OpenAPI operationId.
func xos15FuncNameOpID(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	opIDs := g.Lookup["OpenAPI.operationId"]
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		if opIDs[fn.Name] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fn.FileName,
			Line:    fn.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XOS-15] SSaC func " + fn.Name + " has no matching OpenAPI operationId",
			Advice:  "Align the SSaC function name with the OpenAPI operationId (operationId: " + fn.Name + ")",
		})
	}
	return diags
}
