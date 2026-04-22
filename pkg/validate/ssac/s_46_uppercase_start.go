//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-46 — Result 타입 대문자 시작

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s46UppercaseStart validates S-46: result type must start with uppercase.
func s46UppercaseStart(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Type == "" {
				continue
			}
			t := stripTypePrefix(seq.Result.Type)
			if t == "" {
				continue
			}
			if goPrimitiveTypes[t] {
				continue
			}
			c := t[0]
			if c >= 'A' && c <= 'Z' {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-46] result type %q must start with uppercase", t),
				Advice:  "Result 타입은 대문자로 시작해야 합니다 (PascalCase)",
			})
		}
	}
	return diags
}
