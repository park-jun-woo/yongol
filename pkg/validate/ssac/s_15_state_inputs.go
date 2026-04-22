//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-15 — @state requires an Inputs field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s15StateInputs validates S-15.
func s15StateInputs(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "state" {
				continue
			}
			if len(seq.Inputs) == 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-15] @state requires Inputs",
					Advice:  "Add an Inputs field to the @state sequence",
				})
			}
		}
	}
	return diags
}
