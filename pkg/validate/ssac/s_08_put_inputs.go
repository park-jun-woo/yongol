//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-8 — @put requires an Inputs field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s08PutInputs validates S-8.
func s08PutInputs(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "put" {
				continue
			}
			if len(seq.Args) == 0 && len(seq.Inputs) == 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-8] @put requires Inputs",
					Advice:  "Add an Inputs field to the @put sequence",
				})
			}
		}
	}
	return diags
}
