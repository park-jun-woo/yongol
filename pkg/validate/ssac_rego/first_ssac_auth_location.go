//ff:func feature=validate type=util control=iteration dimension=2 topic=policy-check
//ff:what firstSSaCAuthLocation — (action, resource) 쌍 → 가장 먼저 등장한 SSaC @auth 위치 매핑

package ssac_rego

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// firstSSaCAuthLocation returns the first source location (file + line) at
// which each (action, resource) pair appears as a SSaC @auth sequence. Used
// to anchor diagnostics to the originating call site.
func firstSSaCAuthLocation(funcs []ssac.ServiceFunc) map[[2]string]PairLocation {
	locs := make(map[[2]string]PairLocation)
	for _, fn := range funcs {
		scanSSaCAuthLocations(fn, locs)
	}
	return locs
}
