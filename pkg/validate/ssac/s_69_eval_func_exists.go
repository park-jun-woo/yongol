//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-69 — @eval Func 은 Func Spec(또는 빌트인) 에 존재해야 함

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s69EvalFuncExists validates S-69: @eval Model must resolve to a known
// FuncSpec — either a project func or a yongol-pkg built-in. Mirrors XFS-39
// for @call but stays inside the SSaC package since @eval signature/STATUS
// rules already live here.
func s69EvalFuncExists(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != parsessac.SeqEval || seq.Model == "" {
				continue
			}
			if lookupEvalSpec(seq.Model, fs.ProjectFuncSpecs, fs.YongolPkgSpecs) != nil {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[S-69] @eval references unknown func " + seq.Model,
				Advice:  "Define a predicate Func at pkg/<package>/<method>.go with @func camelCaseName and bool return.",
			})
		}
	}
	return diags
}
