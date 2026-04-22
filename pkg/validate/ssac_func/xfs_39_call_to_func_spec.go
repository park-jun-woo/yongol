//ff:func feature=validate type=rule control=iteration dimension=2 topic=func-check
//ff:what XFS-39 — @call → func 구현 존재

package ssac_func

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs39CallToFuncSpec validates XFS-39: @call → func 구현 존재
func xfs39CallToFuncSpec(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	specs := g.Lookup["Func.spec"]
	if specs == nil {
		specs = map[string]bool{}
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" || seq.Model == "" {
				continue
			}
			if !strings.Contains(seq.Model, ".") {
				continue
			}
			key := normalizedCallKey(seq.Model)
			if !specs[key] {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[XFS-39] @call references non-existent func spec " + seq.Model,
					Advice:  "pkg/ 아래 함수 " + seq.Model + " 를 정의하세요",
				})
			}
		}
	}
	return diags
}
