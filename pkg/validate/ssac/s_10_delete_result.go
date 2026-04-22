//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-10 — @delete Result 부재

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s10DeleteResult validates S-10.
func s10DeleteResult(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "delete" {
				continue
			}
			if seq.Result != nil {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-10] @delete must not have Result",
					Advice:  "@delete 시퀀스에서 Result 항목을 제거하세요",
				})
			}
		}
	}
	return diags
}
