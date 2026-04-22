//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-40 — the @subscribe parameter must be named 'message' and be the sole parameter

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s40SubscribeSingleParam validates S-40: @subscribe function's parameter
// variable must be a single one named "message".
func s40SubscribeSingleParam(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil {
			continue
		}
		if fn.Param == nil || fn.Param.VarName != "message" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    fn.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[S-40] @subscribe parameter variable must be named 'message'",
				Advice:  "Declare the @subscribe function parameter as a single variable named 'message'",
			})
		}
	}
	return diags
}
