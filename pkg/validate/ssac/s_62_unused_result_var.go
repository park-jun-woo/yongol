//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-62 — ERROR when a declared result variable is never referenced by a subsequent sequence

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s62UnusedResultVar validates S-62: a result variable declared in a sequence
// must be referenced at least once in the sequences that follow it (Inputs
// values, Fields values, or Target). Emits an ERROR when never referenced.
func s62UnusedResultVar(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		seqs := fn.Sequences
		for i, seq := range seqs {
			if seq.Result == nil || seq.Result.Var == "" {
				continue
			}
			varName := seq.Result.Var
			if varName == "_" {
				continue
			}
			if s62unusedInSubsequent(varName, seqs, i+1) {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-62] result variable %q is never used in a subsequent sequence", varName),
					Advice:  "Remove the unused variable, or reference it in @response or a later sequence",
				})
			}
		}
	}
	return diags
}
