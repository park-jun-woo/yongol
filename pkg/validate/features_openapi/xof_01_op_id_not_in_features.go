//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-openapi
//ff:what XOF-01 — OpenAPI operationId가 features에 없으면 ERROR

package features_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xof01OpIDNotInFeatures validates XOF-01: every OpenAPI operationId must
// be listed in features.yaml.
func xof01OpIDNotInFeatures(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil || fs.Features == nil {
		return nil
	}
	featureOps := make(map[string]bool, len(fs.Features))
	for _, f := range fs.Features {
		featureOps[f.Op] = true
	}
	var diags []diagnostic.Diagnostic
	for opID := range g.Lookup["OpenAPI.operationId"] {
		if featureOps[opID] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        "api/openapi.yaml",
			Line:        fs.OpenAPILines.OperationLine(opID),
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     "[XOF-01] OpenAPI operationId " + opID + " is not listed in features.yaml",
			Advice:      "Add this operationId to features.yaml",
			OperationID: opID,
		})
	}
	return diags
}
