//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-30 — variables referenced in @response Fields must be declared

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s30MessageDeclared validates S-30: @response Fields' leading variable must be declared.
func s30MessageDeclared(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for i, seq := range fn.Sequences {
			if len(seq.Fields) == 0 {
				continue
			}
			declared := declaredVars(fn, i)
			for _, val := range seq.Fields {
				ref := inputValueRefBase(val)
				if ref == "" || isImplicitVar(ref) || declared[ref] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-30] @response field variable %q used before declaration", ref),
					Advice:  fmt.Sprintf("Declare variable %q as the result of a preceding @get/@post sequence", ref),
				})
			}
		}
	}
	return diags
}
