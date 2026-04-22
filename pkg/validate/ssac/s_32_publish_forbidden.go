//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-32 — @publish에서 query 금지

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s32PublishForbidden validates S-32: @publish Inputs may not contain query.
func s32PublishForbidden(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "publish" {
				continue
			}
			for _, val := range seq.Inputs {
				if val == "query" {
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: "[S-32] @publish cannot use query",
						Advice:  "@publish 인자에서 query.* 참조를 제거하세요",
					})
				}
			}
		}
	}
	return diags
}
