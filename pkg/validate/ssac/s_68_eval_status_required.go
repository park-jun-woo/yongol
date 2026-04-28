//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-68 — @eval STATUS 명시 필수 (디폴트 없음)

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s68EvalStatusRequired validates S-68: @eval has no implicit STATUS default;
// authors must supply one explicitly so the early-return contract is visible.
// (S-58 then validates the value falls in 100-599.)
func s68EvalStatusRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != parsessac.SeqEval {
				continue
			}
			if seq.ErrStatus != 0 {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[S-68] @eval requires an explicit STATUS code",
				Advice:  `Append a status: @eval pkg.Func({...}) "msg" 402`,
			})
		}
	}
	return diags
}
