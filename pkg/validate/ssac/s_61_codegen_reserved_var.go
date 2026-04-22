//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-61 — prevents result variable names from colliding with codegen-reserved identifiers

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s61CodegenReservedVar validates S-61: result variable must not collide with
// codegen-reserved names (s, ctx, err, tx, qtx, db, api, conn, srv, r).
// These names are used by the generated Go code and cause compilation errors.
func s61CodegenReservedVar(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Var == "" {
				continue
			}
			v := seq.Result.Var
			if !codegenReservedVars[v] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-61] variable name %q is a codegen-reserved identifier (used for Server receiver, context, error, etc.)", v),
				Advice:  "Rename it to something descriptive (e.g. s → summary, ctx → callCtx)",
			})
		}
	}
	return diags
}
