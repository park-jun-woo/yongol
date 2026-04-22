//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-5 — @post Inputs 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s05PostInputs validates S-5.
func s05PostInputs(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "post" {
				continue
			}
			if len(seq.Args) == 0 && len(seq.Inputs) == 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-5] @post requires Inputs",
					Advice:  "@post 시퀀스에 Inputs 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
