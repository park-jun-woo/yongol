//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-5 — @post requires an Inputs field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s05PostInputs validates S-5.
func s05PostInputs(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "post" {
				continue
			}
			if len(seq.Args) == 0 && len(seq.Inputs) == 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-5] @post requires Inputs",
					Advice:  "Add an Inputs field to the @post sequence",
				})
			}
		}
	}
	return diags
}
