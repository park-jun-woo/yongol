//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-18 — @auth requires an Action field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s18AuthAction validates S-18.
func s18AuthAction(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "auth" {
				continue
			}
			if seq.Action == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-18] @auth requires Action",
					Advice:  "Add an Action field to the @auth sequence (e.g. read, write)",
				})
			}
		}
	}
	return diags
}
