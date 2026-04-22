//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-41 — currentUser is forbidden inside @subscribe functions

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s41SubscribeNoCurrentUser validates S-41: @subscribe Args may not use currentUser.
func s41SubscribeNoCurrentUser(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil {
			continue
		}
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source == "currentUser" {
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: "[S-41] @subscribe cannot use currentUser",
						Advice:  "@subscribe functions cannot use HTTP inputs (request, query, or currentUser)",
					})
				}
			}
		}
	}
	return diags
}
