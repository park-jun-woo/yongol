//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-34 — Go reserved words are forbidden as result variable names

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s34GoReservedWord validates S-34: result variable must not be a Go reserved word.
func s34GoReservedWord(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Var == "" {
				continue
			}
			if !goReservedWords[seq.Result.Var] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-34] Go reserved word %q used as variable name", seq.Result.Var),
				Advice:  fmt.Sprintf("Go reserved word %q cannot be used as a variable name", seq.Result.Var),
			})
		}
	}
	return diags
}
