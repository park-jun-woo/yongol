//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what XSS-47 — @call arg source 미정의 (WARNING)

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss47CallSourceVarUndefined validates XSS-47: each @call arg's source variable
// must be declared earlier in the same function (implicit reserved names skipped).
func xss47CallSourceVarUndefined(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for i, seq := range fn.Sequences {
			if seq.Type != "call" {
				continue
			}
			declared := declaredVars(fn, i)
			for _, arg := range seq.Args {
				name := arg.Source
				if name == "" || isImplicitVar(name) || declared[name] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: fmt.Sprintf("[XSS-47] @call arg source %q is undefined", name),
					Advice:  fmt.Sprintf("@call 인자 source 변수 %q 를 미리 선언하세요", name),
				})
			}
		}
	}
	return diags
}
