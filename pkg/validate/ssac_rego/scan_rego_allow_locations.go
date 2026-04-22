//ff:func feature=validate type=util control=iteration dimension=2 topic=policy-check
//ff:what scanRegoAllowLocations — 단일 Rego Policy 내 allow 규칙 위치를 locs 에 기록

package ssac_rego

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

// scanRegoAllowLocations iterates over one policy's rules and registers the
// first occurrence location for each (action, resource) pair into locs.
func scanRegoAllowLocations(p rego.Policy, locs map[[2]string]PairLocation) {
	for _, r := range p.Rules {
		for _, action := range r.Actions {
			key := [2]string{action, r.Resource}
			if _, ok := locs[key]; ok {
				continue
			}
			locs[key] = PairLocation{File: p.File, Line: r.SourceLine}
		}
	}
}
