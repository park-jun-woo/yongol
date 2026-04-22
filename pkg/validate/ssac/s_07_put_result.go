//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-7 — @put must not have a Result field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s07PutResult validates S-7.
func s07PutResult(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "put" {
				continue
			}
			if seq.Result != nil {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-7] @put must not have Result",
					Advice:  "Remove the Result field from the @put sequence (re-fetch with @get if the value is needed)",
				})
			}
		}
	}
	return diags
}
