//ff:func feature=policy type=util control=iteration dimension=1
//ff:what collectOpaErrors — ast.Errors → R-1 진단 배열로 변환

package rego

import (
	"github.com/open-policy-agent/opa/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// collectOpaErrors converts each entry of ast.Errors into an R-1 parse
// diagnostic preserving line info when available.
func collectOpaErrors(path string, errs ast.Errors) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, e := range errs {
		line := 0
		if e.Location != nil {
			line = e.Location.Row
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: "R-1: Rego parse error: " + e.Message + " → 권고: OPA v1 문법 (allow if { … }) 에 맞게 수정하세요",
		})
	}
	return diags
}
