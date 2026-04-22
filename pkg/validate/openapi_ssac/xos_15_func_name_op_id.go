//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-15 — SSaC funcName이 OpenAPI operationId에 정의되어 있는지 검증

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
			Advice:  "SSaC 함수명을 OpenAPI operationId 와 일치시키세요 (operationId: " + fn.Name + ")",
		})
	}
	return diags
}
