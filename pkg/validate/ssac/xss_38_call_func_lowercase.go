//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what XSS-38 — @call function name must not start with a lowercase letter

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss38CallFuncLowercase validates XSS-38: @call's Method name must start
// with an uppercase letter (Go-exported). Lowercase starts mean unexported.
func xss38CallFuncLowercase(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" {
				continue
			}
			method := extractMethod(seq)
			if method == "" {
				continue
			}
			c := method[0]
			if c >= 'A' && c <= 'Z' {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[XSS-38] @call function %q must start with uppercase", method),
				Advice:  "@call function names must start with an uppercase letter (Go exported)",
			})
		}
	}
	return diags
}
