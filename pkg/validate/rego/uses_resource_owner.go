//ff:func feature=validate type=util control=iteration dimension=1 topic=rego-structural
//ff:what usesResourceOwner — Rego AllowRule 중 UsesOwner가 하나라도 있는지

package rego

import regoparser "github.com/park-jun-woo/yongol/pkg/parser/rego"

func usesResourceOwner(rules []regoparser.AllowRule) bool {
	for _, r := range rules {
		if r.UsesOwner {
			return true
		}
	}
	return false
}
