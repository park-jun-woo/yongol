//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-structural
//ff:what O-3 — path 템플릿 변수와 parameters[].name 일치 검증

package openapi

import (
	"regexp"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// rePathVar captures the variable name inside an OpenAPI path template
// segment (e.g. `/workflows/{id}/actions` → ["id"]). Used by O-3 and helper
// funcs in the same package.
var rePathVar = regexp.MustCompile(`\{([^{}/]+)\}`)

// o03PathTemplateParam validates O-3: every `{X}` in a path template must
// correspond to a parameter declared with `in: path, name: X` either at the
// path-item level or the operation level. OpenAPI 3.x mandates this, but the
// kin-openapi loader accepts mismatches silently — this rule makes the
// mismatch an explicit ERROR before it propagates to SSaC / STML / codegen.
func o03PathTemplateParam(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for path, item := range fs.OpenAPIDoc.Paths.Map() {
		if item == nil {
			continue
		}
		want := collectPathVars(path)
		for method, op := range item.Operations() {
			got := collectDeclaredPathParams(item.Parameters, op.Parameters)
			line := fs.OpenAPILines.OperationLine(op.OperationID)
			if line == 0 {
				line = fs.OpenAPILines.PathLine(path)
			}
			diags = append(diags, comparePathVars(path, method, line, want, got)...)
		}
	}
	return diags
}
