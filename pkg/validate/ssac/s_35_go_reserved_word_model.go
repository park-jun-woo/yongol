//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-35 — Go 예약어 result type 금지

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s35GoReservedWordModel validates S-35: result type must not be a Go reserved word.
func s35GoReservedWordModel(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Type == "" {
				continue
			}
			t := stripTypePrefix(seq.Result.Type)
			if !goReservedWords[t] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-35] result type %q is a Go reserved word", t),
				Advice:  fmt.Sprintf("Go 예약어 %q 는 타입명으로 사용할 수 없습니다", t),
			})
		}
	}
	return diags
}
