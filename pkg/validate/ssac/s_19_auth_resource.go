//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-19 — @auth Resource 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s19AuthResource validates S-19.
func s19AuthResource(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "auth" {
				continue
			}
			if seq.Resource == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-19] @auth requires Resource",
					Advice:  "@auth 시퀀스에 Resource 항목을 추가하세요",
				})
			}
		}
	}
	return diags
}
