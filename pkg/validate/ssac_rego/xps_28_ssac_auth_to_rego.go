//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XPS-28 — SSaC @auth → Rego allow

package ssac_rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xps28SSaCAuthToRego validates XPS-28: every SSaC @auth (action, resource)
// pair has a matching Rego allow rule. Missing coverage means runtime requests
// will always be denied.
func xps28SSaCAuthToRego(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}

	regoPairs := collectRegoAllowPairs(fs)
	ssacPairs := collectSSaCAuthPairs(fs.ServiceFuncs)
	pairLoc := firstSSaCAuthLocation(fs.ServiceFuncs)

	var diags []diagnostic.Diagnostic
	for pair := range ssacPairs {
		if d, ok := xps28MissingRegoDiag(pair, regoPairs, pairLoc); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
