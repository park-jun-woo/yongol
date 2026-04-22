//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-20 — @auth Message 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s20AuthMessage validates S-20.
func s20AuthMessage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "auth" {
				continue
			}
			if seq.Message == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-20] @auth requires Message",
					Advice:  "@auth 시퀀스에 Message 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
