//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-21 — @call Model 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s21CallModel validates S-21.
func s21CallModel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" {
				continue
			}
			if seq.Model == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-21] @call requires Model (package.Func)",
					Advice:  "@call 시퀀스에 Model 항목을 추가하세요 (package.Func 형식)",
				})
			}
		}
	}
	return diags
}
