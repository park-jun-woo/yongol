//ff:func feature=validate type=util control=iteration dimension=2 topic=policy-check
//ff:what collectSSaCAuthPairs — SSaC @auth 시퀀스에서 (action, resource) 쌍 수집

package ssac_rego

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// collectSSaCAuthPairs gathers (action, resource) pairs from every @auth
// sequence across the provided SSaC service functions.
func collectSSaCAuthPairs(funcs []ssac.ServiceFunc) map[[2]string]bool {
	pairs := make(map[[2]string]bool)
	for _, fn := range funcs {
		for _, seq := range fn.Sequences {
			if seq.Type != "auth" {
				continue
			}
			pairs[[2]string{seq.Action, seq.Resource}] = true
		}
	}
	return pairs
}
