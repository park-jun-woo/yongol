//ff:func feature=validate type=rule control=iteration dimension=2 topic=func-check
//ff:what XFS-63 — @call Func signature must return (Response, error)

package ssac_func

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xfs63CallFuncSignature(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" || seq.Model == "" {
				continue
			}
			if !strings.Contains(seq.Model, ".") {
				continue
			}
			spec := findFuncSpec(normalizedCallKey(seq.Model), fs.ProjectFuncSpecs, fs.YongolPkgSpecs)
			if spec == nil {
				continue
			}
			if len(spec.ReturnTypes) == 2 && spec.ReturnTypes[1] == "error" {
				continue
			}
			actual := "(no return)"
			if len(spec.ReturnTypes) > 0 {
				actual = "(" + joinReturnTypes(spec.ReturnTypes) + ")"
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:        fn.FileName,
				Line:        seq.Line,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     fmt.Sprintf("[XFS-63] @call %s signature must return (Response, error), got %s", seq.Model, actual),
				Advice:      fmt.Sprintf("func %s(req T) (Response, error) 형태로 수정하세요. side-effect 전용이면 빈 Response struct 를 사용하세요.", callFuncName(seq.Model)),
				OperationID: fn.Name,
			})
		}
	}
	return diags
}

