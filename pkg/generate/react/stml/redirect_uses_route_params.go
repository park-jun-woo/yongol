//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ActionBlock의 data-redirect-params 소스에 route.* 가 있는지 판별한다 (useParams import 결정)
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// redirectUsesRouteParams reports whether any of the action's redirect
// param sources is a route.<Name> reference, which requires the
// useParams() hook (page-flow Phase008 — the data-redirect twin of
// linkUsesRouteParams).
func redirectUsesRouteParams(a stmlparser.ActionBlock) bool {
	for _, p := range a.RedirectParams {
		if strings.HasPrefix(p.Source, "route.") {
			return true
		}
	}
	return false
}
