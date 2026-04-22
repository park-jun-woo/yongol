//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateRego — Rego Policy에서 auth 쌍, ownership, claims, roles 추출
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateRego(g *rule.Ground, fs *yongol.Fullstack) {
	authPairs := make(rule.StringSet)
	claimsRefs := make(rule.StringSet)
	regoRoles := make(rule.StringSet)

	for _, p := range fs.ParsedPolicies {
		populateRegoPolicy(p, authPairs, claimsRefs, regoRoles)
	}
	g.Pairs["Policy.auth"] = authPairs
	g.Lookup["Rego.claims"] = claimsRefs
	g.Lookup["Rego.roles"] = regoRoles
}
