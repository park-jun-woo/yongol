//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-35 — Go reserved words are forbidden as result type names

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s35GoReservedWordModel validates S-35: result type must not be a Go reserved word.
func s35GoReservedWordModel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Type == "" {
				continue
			}
			t := stripTypePrefix(seq.Result.Type)
			if !goReservedWords[t] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-35] result type %q is a Go reserved word", t),
				Advice:  fmt.Sprintf("Go reserved word %q cannot be used as a type name", t),
			})
		}
	}
	return diags
}
