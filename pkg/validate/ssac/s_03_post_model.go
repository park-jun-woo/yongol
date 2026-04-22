//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-3 — @post requires a Model field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s03PostModel validates S-3.
func s03PostModel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "post" {
				continue
			}
			if seq.Model == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-3] @post requires Model",
					Advice:  "Add a Model field to the @post sequence",
				})
			}
		}
	}
	return diags
}
