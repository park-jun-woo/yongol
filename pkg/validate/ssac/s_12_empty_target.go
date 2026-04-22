//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-12 — @empty Target 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s12EmptyTarget validates S-12.
func s12EmptyTarget(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "empty" {
				continue
			}
			if seq.Target == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-12] @empty requires Target",
					Advice:  "@empty 시퀀스에 Target 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
