//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-45 — @response is forbidden inside @subscribe functions

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s45SubscribeNoResponse validates S-45: @subscribe funcs cannot use @response.
func s45SubscribeNoResponse(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil {
			continue
		}
		for _, seq := range fn.Sequences {
			if seq.Type != "response" {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[S-45] @subscribe cannot use @response",
				Advice:  "Remove the @response sequence from the @subscribe function (responses are HTTP-only)",
			})
		}
	}
	return diags
}
