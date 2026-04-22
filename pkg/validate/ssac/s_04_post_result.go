//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-4 — @post Result 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s04PostResult validates S-4.
func s04PostResult(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "post" {
				continue
			}
			if seq.Result == nil {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-4] @post requires Result",
					Advice:  "@post 시퀀스에 Result 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
