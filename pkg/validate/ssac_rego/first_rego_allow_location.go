//ff:func feature=validate type=util control=iteration dimension=2 topic=policy-check
//ff:what firstRegoAllowLocation — (action, resource) 쌍 → 가장 먼저 등장한 Rego allow 위치 매핑

package ssac_rego

import (
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

// firstRegoAllowLocation returns the first source location (file + line) at
// which each (action, resource) pair appears as a Rego allow rule.
func firstRegoAllowLocation(policies []rego.Policy) map[[2]string]PairLocation {
	locs := make(map[[2]string]PairLocation)
	for _, p := range policies {
		scanRegoAllowLocations(p, locs)
	}
	return locs
}
