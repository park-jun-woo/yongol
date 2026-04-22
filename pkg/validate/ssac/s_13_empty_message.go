//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-13 — @empty Message 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s13EmptyMessage validates S-13.
func s13EmptyMessage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "empty" {
				continue
			}
			if seq.Message == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-13] @empty requires Message",
					Advice:  "@empty 시퀀스에 Message 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
