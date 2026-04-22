//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-44 — the message variable is forbidden in HTTP handler functions

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s44HTTPForbiddenMessage validates S-44: HTTP (non-subscribe) funcs may not use message.
func s44HTTPForbiddenMessage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source == "message" {
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: "[S-44] HTTP function cannot use message",
						Advice:  "HTTP functions cannot use the message variable (message is exclusive to @subscribe)",
					})
				}
			}
		}
	}
	return diags
}
