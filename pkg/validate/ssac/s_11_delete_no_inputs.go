//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-11 — @delete Inputs 없음 WARNING

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s11DeleteNoInputs validates S-11: @delete with no inputs is a warning.
func s11DeleteNoInputs(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "delete" {
				continue
			}
			if len(seq.Args) > 0 || len(seq.Inputs) > 0 {
				continue
			}
			if seq.SuppressWarn {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: "[S-11] @delete has no inputs (all rows may be affected)",
				Advice:  "@delete 시퀀스에 식별 조건을 Inputs 로 추가하거나 의도된 경우 // ff:allow-empty-delete 주석으로 명시하세요",
			})
		}
	}
	return diags
}
