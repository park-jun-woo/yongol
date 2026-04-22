//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-17 — @state Message 필수

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
					Advice:  "@state 시퀀스에 Message 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
