//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-67 — @eval Func 시그니처는 func(req T) bool 이어야 함

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s67EvalFuncSignature validates S-67: every @eval target Func must declare
// a single `bool` return value. Funcs returning `error` belong to @call;
// multi-value returns and other shapes are rejected outright.
func s67EvalFuncSignature(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != parsessac.SeqEval || seq.Model == "" {
				continue
			}
			spec := lookupEvalSpec(seq.Model, fs.ProjectFuncSpecs, fs.YongolPkgSpecs)
			if spec == nil {
				// S-69 reports missing func spec; signature check defers.
				continue
			}
			if isBoolPredicateSignature(spec) {
				continue
			}
			ret := "(no return)"
			if len(spec.ReturnTypes) > 0 {
				ret = "(" + joinTypes(spec.ReturnTypes) + ")"
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-67] @eval %s must return a single bool; got %s", seq.Model, ret),
				Advice:  "Predicate funcs return bool. Use @call for funcs that return error.",
			})
		}
	}
	return diags
}
