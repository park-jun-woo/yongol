//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-33 — ssac 예약 소스를 result 변수로 사용 금지

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s33ReservedSource validates S-33: result variable name must not collide with
// reserved source names (request, currentUser, config, query, message).
func s33ReservedSource(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Var == "" {
				continue
			}
			if !reservedSourceNames[seq.Result.Var] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-33] reserved source %q used as result variable", seq.Result.Var),
				Advice:  fmt.Sprintf("변수명 %q 는 예약어이므로 다른 이름을 사용하세요", seq.Result.Var),
			})
		}
	}
	return diags
}
