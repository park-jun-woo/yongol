//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-24 — @publish requires a Payload field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s24PublishPayload validates S-24.
func s24PublishPayload(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "publish" {
				continue
			}
			if len(seq.Inputs) == 0 && len(seq.Fields) == 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-24] @publish requires Payload",
					Advice:  "Add a Payload field to the @publish sequence",
				})
			}
		}
	}
	return diags
}
