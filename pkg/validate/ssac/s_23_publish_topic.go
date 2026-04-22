//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-23 — @publish requires a Topic field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s23PublishTopic validates S-23.
func s23PublishTopic(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "publish" {
				continue
			}
			if seq.Topic == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-23] @publish requires Topic",
					Advice:  "Add a Topic field to the @publish sequence",
				})
			}
		}
	}
	return diags
}
