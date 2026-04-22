//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-27 — variables must be declared before use in Args

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s27VarDeclared validates S-27: every variable referenced by Args.Source
// must be a previously declared result variable (or an implicit reserved name).
func s27VarDeclared(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for i, seq := range fn.Sequences {
			declared := declaredVars(fn, i)
			for _, arg := range seq.Args {
				name := arg.Source
				if name == "" || isImplicitVar(name) || declared[name] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-27] variable %q used before declaration", name),
					Advice:  fmt.Sprintf("Declare variable %q as the result of a preceding @get/@post sequence", name),
				})
			}
		}
	}
	return diags
}
