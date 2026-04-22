//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-16 — @state requires a Transition field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s16StateTransition validates S-16.
func s16StateTransition(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "state" {
				continue
			}
			if seq.Transition == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-16] @state requires Transition",
					Advice:  "Add a Transition field to the @state sequence (e.g. Draft-->Published)",
				})
			}
		}
	}
	return diags
}
