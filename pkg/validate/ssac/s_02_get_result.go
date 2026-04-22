//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-2 — @get Result 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s02GetResult validates S-2: @get requires Result
func s02GetResult(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "get" {
				continue
			}
			if seq.Result == nil {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-2] @get requires Result",
					Advice:  "@get 시퀀스에 Result 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
