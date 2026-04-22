//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-9 — @delete requires a Model field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s09DeleteModel validates S-9.
func s09DeleteModel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "delete" {
				continue
			}
			if seq.Model == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-9] @delete requires Model",
					Advice:  "Add a Model field to the @delete sequence",
				})
			}
		}
	}
	return diags
}
