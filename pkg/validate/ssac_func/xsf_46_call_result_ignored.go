//ff:func feature=validate type=rule control=iteration dimension=2 topic=func-check
//ff:what XSF-46 — @result 없지만 func Response 있음

package ssac_func

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsf46CallResultIgnored validates XSF-46: @result 없지만 func Response 있음.
// WARNING level — caller silently drops a return value the callee provides.
func xsf46CallResultIgnored(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" || seq.Model == "" || seq.Result != nil {
				continue
			}
			if !strings.Contains(seq.Model, ".") {
				continue
			}
			spec := findFuncSpec(normalizedCallKey(seq.Model), fs.ProjectFuncSpecs, fs.YongolPkgSpecs)
			if spec == nil {
				continue
			}
			if len(spec.ResponseFields) > 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: "[XSF-46] @call " + seq.Model + " ignores Response fields",
					Advice:  "func " + seq.Model + " 의 반환값이 무시됩니다 — @get 결과 변수에 바인드하세요",
				})
			}
		}
	}
	return diags
}
