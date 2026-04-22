//ff:func feature=validate type=rule control=iteration dimension=2 topic=policy-check
//ff:what XSP-29 — Rego allow → SSaC @auth

package ssac_rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsp29RegoAllowToSSaC validates XSP-29: every Rego allow (action, resource)
// pair has a matching SSaC @auth sequence. Unmatched rules are unused policy.
func xsp29RegoAllowToSSaC(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}

	regoPairs := collectRegoAllowPairs(fs)
	ssacPairs := collectSSaCAuthPairs(fs.ServiceFuncs)
	pairLoc := firstRegoAllowLocation(fs.ParsedPolicies)

	var diags []diagnostic.Diagnostic
	for pair := range regoPairs {
		if d, ok := xsp29MissingSSaCDiag(pair, ssacPairs, pairLoc); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
