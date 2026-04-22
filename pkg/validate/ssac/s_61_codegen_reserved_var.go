//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-61 — 코드젠 예약 변수명 충돌 방지

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s61CodegenReservedVar validates S-61: result variable must not collide with
// codegen-reserved names (s, ctx, err, tx, qtx, db, api, conn, srv, r).
// These names are used by the generated Go code and cause compilation errors.
func s61CodegenReservedVar(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Var == "" {
				continue
			}
			v := seq.Result.Var
			if !codegenReservedVars[v] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-61] 변수명 %q 는 코드젠 예약어입니다 (Server receiver, context, error 등에서 사용)", v),
				Advice:  "의미를 드러내는 이름으로 바꾸세요 (예: s → summary, ctx → callCtx)",
			})
		}
	}
	return diags
}
