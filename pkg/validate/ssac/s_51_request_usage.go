//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-51 — OpenAPI request field가 SSaC에서 사용

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s51RequestUsage validates S-51: every OpenAPI request field must be used by
// the SSaC function (warns when an OpenAPI field is declared but never referenced).
func s51RequestUsage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		expected, ok := g.Lookup["OpenAPI.request."+fn.Name]
		if !ok {
			continue
		}
		used := map[string]bool{}
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source == "request" && arg.Field != "" {
					used[arg.Field] = true
				}
			}
			// CRUD/@call sequences carry inputs in seq.Inputs (key → "source.field").
			for _, val := range seq.Inputs {
				v := strings.TrimSpace(val)
				if !strings.HasPrefix(v, "request.") {
					continue
				}
				field := strings.TrimPrefix(v, "request.")
				if field == "" {
					continue
				}
				used[field] = true
			}
			// @verify-password stores its two request-derived expressions outside
			// seq.Inputs — mark them used if they originate from `request.`.
			for _, expr := range []string{seq.EmailExpr, seq.PasswordExpr} {
				e := strings.TrimSpace(expr)
				if !strings.HasPrefix(e, "request.") {
					continue
				}
				field := strings.TrimPrefix(e, "request.")
				if field == "" {
					continue
				}
				used[field] = true
			}
		}
		for field := range expected {
			if used[field] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    fn.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[S-51] OpenAPI request field %q declared but unused in SSaC", field),
				Advice:  "OpenAPI 에 선언했지만 미사용. 제거하거나 SSaC 에서 사용하세요",
			})
		}
	}
	return diags
}
