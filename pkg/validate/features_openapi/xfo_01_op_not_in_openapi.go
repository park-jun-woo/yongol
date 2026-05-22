//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-openapi
//ff:what XFO-01 — features op이 OpenAPI operationId에 없으면 ERROR

package features_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfo01OpNotInOpenAPI validates XFO-01: every features op must have a
// matching OpenAPI operationId.
func xfo01OpNotInOpenAPI(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil || fs.Features == nil {
		return nil
	}
	opIDs := g.Lookup["OpenAPI.operationId"]
	var diags []diagnostic.Diagnostic
	for _, f := range fs.Features {
		if opIDs[f.Op] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        "features.yaml",
			Line:        f.Line,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     "[XFO-01] features op " + f.Op + " has no matching OpenAPI operationId",
			Advice:      "Add operationId: " + f.Op + " to the OpenAPI spec, or remove this feature entry",
			OperationID: f.Op,
		})
	}
	return diags
}
