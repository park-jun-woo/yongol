//ff:func feature=validate type=rule control=iteration dimension=2 topic=func-check
//ff:what XFS-45 — @result 있지만 func Response 없음

package ssac_func

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs45CallResultMissing validates XFS-45: @result 있지만 func Response 없음.
// ERROR level — caller tries to bind a variable that the callee never returns.
func xfs45CallResultMissing(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" || seq.Model == "" || seq.Result == nil {
				continue
			}
			if !strings.Contains(seq.Model, ".") {
				continue
			}
			spec := findFuncSpec(normalizedCallKey(seq.Model), fs.ProjectFuncSpecs, fs.YongolPkgSpecs)
			if spec == nil {
				continue
			}
			if len(spec.ResponseFields) == 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[XFS-45] @call " + seq.Model + " binds @result but func has no Response fields",
					Advice:  "@call 결과를 받지 않거나 func 에 ResponseFields 를 정의하세요",
				})
			}
		}
	}
	return diags
}
