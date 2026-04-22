//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-18 — @auth Action 필수

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s18AuthAction validates S-18.
func s18AuthAction(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "auth" {
				continue
			}
			if seq.Action == "" {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-18] @auth requires Action",
					Advice:  "@auth 시퀀스에 Action 항목을 추가하세요 (예: read, write)",
				})
			}
		}
	}
	return diags
}
