//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-11 — WARNING when @delete has no Inputs

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s11DeleteNoInputs validates S-11: @delete with no inputs is a warning.
func s11DeleteNoInputs(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "delete" {
				continue
			}
			if len(seq.Args) > 0 || len(seq.Inputs) > 0 {
				continue
			}
			if seq.SuppressWarn {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: "[S-11] @delete has no inputs (all rows may be affected)",
				Advice:  "Add identifying conditions to the @delete sequence Inputs, or annotate with // ff:allow-empty-delete if intentional",
			})
		}
	}
	return diags
}
