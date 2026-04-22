//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-6 — @put requires a Model field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s06PutModel validates S-6.
func s06PutModel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "put" {
				continue
			}
			if seq.Model == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-6] @put requires Model",
					Advice:  "Add a Model field to the @put sequence",
				})
			}
		}
	}
	return diags
}
