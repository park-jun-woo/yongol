//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-42 — request inputs are forbidden inside @subscribe functions

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s42SubscribeForbiddenRequest validates S-42: @subscribe Args may not use request.
func s42SubscribeForbiddenRequest(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil {
			continue
		}
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source == "request" {
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: "[S-42] @subscribe cannot use request",
						Advice:  "@subscribe functions cannot use HTTP inputs (request, query, or currentUser)",
					})
				}
			}
		}
	}
	return diags
}
