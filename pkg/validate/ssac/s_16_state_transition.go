//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-16 — @state Transition 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s16StateTransition validates S-16.
func s16StateTransition(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "state" {
				continue
			}
			if seq.Transition == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-16] @state requires Transition",
					Advice:  "@state 시퀀스에 Transition 항목을 추가하세요 (예: Draft-->Published)",
				})
			}
		}
	}
	return diags
}
