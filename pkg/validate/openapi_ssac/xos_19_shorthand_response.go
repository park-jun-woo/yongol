//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-19 — verifies that shorthand @response fields are present in the OpenAPI response schema

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos19ShorthandResponse validates XOS-19: when a SSaC func uses shorthand
// @response (@response varName), the resolved JSON field names match the
// OpenAPI 2xx response schema. Page[T]/Cursor[T]/[]T wrappers are skipped.
func xos19ShorthandResponse(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		varName := shorthandResponseTarget(fn)
		if varName == "" {
			continue
		}
		if !varDeclaredInFunc(fn, varName) {
			diags = append(diags, diagnostic.Diagnostic{
				File:        fn.FileName,
				Line:        fn.Line,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     "[XOS-19] shorthand response variable \"" + varName + "\" not declared in " + fn.Name,
				Advice:      "Declare the shorthand @response variable " + varName + " with @get/@call or similar first",
				OperationID: fn.Name,
			})
			continue
		}
		if shorthandWrapperSkip(fn, varName) {
			continue
		}
		fields := g.Schemas["SSaC.response."+fn.Name]
		if len(fields) == 0 {
			continue
		}
		opProps := toSet(g.Schemas["OpenAPI.response."+fn.Name])
		if len(opProps) == 0 {
			continue
		}
		diags = append(diags, collectMissingProps(fn, fields, opProps, "XOS-19")...)
	}
	return diags
}
