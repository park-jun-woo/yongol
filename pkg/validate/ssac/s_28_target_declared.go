//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-28 — @empty Target 변수 선언

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s28TargetDeclared validates S-28: Target's leading variable must be declared.
func s28TargetDeclared(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for i, seq := range fn.Sequences {
			if seq.Target == "" {
				continue
			}
			ref := strings.SplitN(seq.Target, ".", 2)[0]
			if ref == "" || isImplicitVar(ref) {
				continue
			}
			declared := declaredVars(fn, i)
			if declared[ref] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-28] Target variable %q used before declaration", ref),
				Advice:  fmt.Sprintf("변수 %q 를 @get/@post 시퀀스 결과로 먼저 선언하세요", ref),
			})
		}
	}
	return diags
}
