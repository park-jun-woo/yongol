//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-31 — config.* 입력 금지

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s31ConfigPrefixForbidden validates S-31: Inputs values may not start with "config.".
func s31ConfigPrefixForbidden(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			for _, val := range seq.Inputs {
				if !strings.HasPrefix(val, "config.") {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-31] config.* input forbidden — use os.Getenv() inside func",
					Advice:  "config.* 참조 대신 직접 입력값으로 받으세요",
				})
			}
		}
	}
	return diags
}
