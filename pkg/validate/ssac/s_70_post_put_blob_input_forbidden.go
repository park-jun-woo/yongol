//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-70 — @post/@put Inputs 의 reserved source 단독 참조 거절 (DDL 쓰기에 blob 박기 차단)

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s70PostPutBlobInputForbidden validates S-70: @post / @put Inputs values
// must not be a standalone reserved source (currentUser, request, query,
// message). Reserved sources must always be referenced in dotted form
// (source.Field). @call is exempt — user-authored Funcs may legitimately
// receive a raw object.
func s70PostPutBlobInputForbidden(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != parsessac.SeqPost && seq.Type != parsessac.SeqPut {
				continue
			}
			for key, val := range seq.Inputs {
				if !isImplicitVar(val) {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf(`[S-70] @%s input %q: reserved source %q must use dotted form (e.g. %s.<Field>); standalone reference forbidden in DDL writes`, seq.Type, key, val, val),
					Advice:  "claim 별 typed 필드를 명시하여 매핑하세요. 통합 객체를 받아야 하면 @call <pkg>.<Func>({...}) 로 우회.",
				})
			}
		}
	}
	return diags
}
