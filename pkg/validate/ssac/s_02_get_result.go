//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-2 — @get requires a Result field

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s02GetResult validates S-2: @get requires Result
func s02GetResult(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "get" {
				continue
			}
			if seq.Result == nil {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-2] @get requires Result",
					Advice:  "Add a Result field to the @get sequence",
				})
			}
		}
	}
	return diags
}
