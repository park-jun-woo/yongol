//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-71 — SSaC Input 값의 변수 prefix 가 해당 시퀀스 지점에서 유효한 scope 에 없으면 ERROR

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func s71UnknownVariable(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for i, seq := range fn.Sequences {
			scope := s71BuildScope(fn, i)
			refs := s71CollectRefs(seq)
			for _, ref := range refs {
				prefix := s71ExtractPrefix(ref)
				if prefix == "" || scope[prefix] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-71] unknown variable %q in @%s input. Available at this point: %s", prefix, seq.Type, s71ScopeList(scope)),
				})
			}
		}
	}
	return diags
}
