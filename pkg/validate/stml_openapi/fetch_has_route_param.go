//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what fetchHasRouteParam — fetch가 route.* 소스의 path 파라미터를 소비하는지 판정

package stml_openapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// fetchHasRouteParam reports whether any data-param-* of the fetch is sourced
// from a route path parameter ("route.<Name>").
func fetchHasRouteParam(f stml.FetchBlock) bool {
	for _, p := range f.Params {
		if strings.HasPrefix(p.Source, "route.") {
			return true
		}
	}
	return false
}
