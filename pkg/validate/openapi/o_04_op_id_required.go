//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-structural
//ff:what O-4 — OpenAPI operation 에 operationId 필수

package openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// o04OpIdRequired validates O-4: every OpenAPI operation under paths.<path>.<method>
// must declare an `operationId`. OpenAPI 3.x treats operationId as optional, but
// yongol uses it as the cross-SSOT linkage key (States, SSaC, STML, Hurl, Rego,
// Func). A missing operationId silently drops the operation from ground Lookup
// tables and causes downstream crosschecks to false-PASS or misreport. This rule
// makes the absence an explicit ERROR before any crosscheck runs.
func o04OpIdRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for path, item := range fs.OpenAPIDoc.Paths.Map() {
		if item == nil {
			continue
		}
		for method, op := range item.Operations() {
			if op == nil || op.OperationID != "" {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    "api/openapi.yaml",
				Line:    fs.OpenAPILines.PathLine(path),
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[O-4] missing operationId at " + method + " " + path,
				Advice:  "operationId: camelCase 형태로 모든 operation 에 명시하세요 (예: GetWorkflow, ExecuteWorkflow)",
			})
		}
	}
	return diags
}
