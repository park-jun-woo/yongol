//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-38 — HTTP-only sources are forbidden in @subscribe Inputs

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s38SubscribeNoHTTPInputs validates S-38: @subscribe Inputs values must not
// reference HTTP-only sources (request, query, currentUser).
func s38SubscribeNoHTTPInputs(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe == nil {
			continue
		}
		for _, seq := range fn.Sequences {
			for _, val := range seq.Inputs {
				ref := strings.SplitN(val, ".", 2)[0]
				switch ref {
				case "request", "query", "currentUser":
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: fmt.Sprintf("[S-38] @subscribe cannot use HTTP input %q", ref),
						Advice:  "@subscribe functions cannot use HTTP inputs (request, query, or currentUser)",
					})
				}
			}
		}
	}
	return diags
}
