//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-12 — @empty requires a Target field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s12EmptyTarget validates S-12.
func s12EmptyTarget(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "empty" {
				continue
			}
			if seq.Target == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-12] @empty requires Target",
					Advice:  "Add a Target field to the @empty sequence",
				})
			}
		}
	}
	return diags
}
