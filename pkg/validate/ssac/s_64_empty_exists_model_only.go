//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-64 — @empty/@exists Target은 Model 변수여야 함 (스칼라 거절)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s64EmptyExistsModelOnly validates S-64: an @empty / @exists guard's Target
// must reference a Model variable (DDL row struct or Func Response struct),
// never a scalar value or a dotted field. Scalar predicates belong in @eval.
//
// Lookups (populated by pkg/ground):
//   - Ground.Types["SSaC.var.<func>.<var>"] = raw type spec (e.g. "Course",
//     "*User", "[]Webhook", "billing.CheckCreditsResponse")
//   - Ground.Lookup["SymbolTable.model"] = set of DDL model PascalCase names
//   - Ground.Types["Struct.<TypeName>.*"] = present iff TypeName is a known
//     Func Response or DDL row struct (set by registerFuncSpec /
//     populateSSaCSymbols)
func s64EmptyExistsModelOnly(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != parsessac.SeqEmpty && seq.Type != parsessac.SeqExists {
				continue
			}
			if seq.Target == "" {
				continue
			}
			// Reject any dotted field access (e.g. "wf.ID", "org.CreditsBalance").
			// Even when the leading variable is a Model, the Target itself must be
			// the variable, not a field of it.
			if strings.Contains(seq.Target, ".") {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-64] @%s target %q must be a Model var (got dotted-field); use @eval for scalar predicates", seq.Type, seq.Target),
					Advice:  "@empty / @exists는 모델 존재 여부 가드입니다. 스칼라 값을 검사하려면 @eval pkg.Func({...}) \"msg\" STATUS 를 사용하세요.",
				})
				continue
			}
			if isImplicitVar(seq.Target) {
				continue
			}
			rawType := g.Types["SSaC.var."+fn.Name+"."+seq.Target]
			if rawType == "" {
				// Variable not declared / unknown type — S-28 / S-27 will report.
				continue
			}
			typeName := stripTypePrefix(rawType)
			if typeName == "" {
				continue
			}
			if isModelType(g, typeName) {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-64] @%s target %q must be a Model var (got %s); use @eval for scalar predicates", seq.Type, seq.Target, rawType),
				Advice:  "@empty / @exists는 모델 존재 여부 가드입니다. 스칼라 값을 검사하려면 @eval pkg.Func({...}) \"msg\" STATUS 를 사용하세요.",
			})
		}
	}
	return diags
}
