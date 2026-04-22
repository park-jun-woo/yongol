//ff:func feature=validate type=util control=iteration dimension=2 topic=policy-check
//ff:what collectRegoAllowPairs — Rego allow 규칙에서 (action, resource) 쌍 수집

package ssac_rego

import "github.com/park-jun-woo/yongol/pkg/yongol"

// collectRegoAllowPairs gathers (action, resource) pairs from every parsed
// Rego allow rule on the Fullstack.
func collectRegoAllowPairs(fs *yongol.Fullstack) map[[2]string]bool {
	pairs := make(map[[2]string]bool)
	if fs == nil {
		return pairs
	}
	for _, p := range fs.ParsedPolicies {
		for _, rule := range p.Rules {
			for _, action := range rule.Actions {
				pairs[[2]string{action, rule.Resource}] = true
			}
		}
	}
	return pairs
}
