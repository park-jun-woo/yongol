//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-15 — @state Inputs 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s15StateInputs validates S-15.
func s15StateInputs(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "state" {
				continue
			}
			if len(seq.Inputs) == 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-15] @state requires Inputs",
					Advice:  "@state 시퀀스에 Inputs 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
