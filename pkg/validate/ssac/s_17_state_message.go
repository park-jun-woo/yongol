//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-17 — @state requires a Message field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s17StateMessage validates S-17.
func s17StateMessage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "state" {
				continue
			}
			if seq.Message == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-17] @state requires Message",
					Advice:  "Add a Message field to the @state sequence",
				})
			}
		}
	}
	return diags
}
