//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-29 — variables used as Inputs values must be declared

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s29InputDeclared validates S-29: Inputs values' leading variable must be declared.
func s29InputDeclared(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for i, seq := range fn.Sequences {
			if len(seq.Inputs) == 0 {
				continue
			}
			declared := declaredVars(fn, i)
			for _, val := range seq.Inputs {
				ref := inputValueRefBase(val)
				if ref == "" || isImplicitVar(ref) || declared[ref] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-29] Input value variable %q used before declaration", ref),
					Advice:  fmt.Sprintf("Declare variable %q as the result of a preceding @get/@post sequence", ref),
				})
			}
		}
	}
	return diags
}
