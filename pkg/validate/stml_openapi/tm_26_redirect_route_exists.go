//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-26 — data-redirect 경로가 어떤 STML 페이지 라우트에도 해석되지 않음 (ERROR, "/"는 인덱스로 허용)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm26RedirectRouteExists checks that an action's data-redirect path
// resolves to some STML page route (explicit data-route or the
// filename-derived patterns, ":param" segments matching any one segment).
// "/" is always allowed — Phase005 emits the index route. A redirect into
// the void would generate a navigate() to a route that renders nothing.
func tm26RedirectRouteExists(a stml.ActionBlock, file string, pages []stml.PageSpec) []diagnostic.Diagnostic {
	if a.Redirect == "" || a.Redirect == "/" {
		return nil
	}
	for _, p := range pages {
		for _, pattern := range stml.RoutePaths(p) {
			if stml.RouteMatchesPath(pattern, a.Redirect) {
				return nil
			}
		}
	}
	return []diagnostic.Diagnostic{{
		File:        file,
		Phase:       diagnostic.PhaseValidate,
		Level:       diagnostic.LevelError,
		Message:     fmt.Sprintf("[TM-26] data-redirect %q on action %q does not resolve to any STML page route", a.Redirect, a.OperationID),
		Advice:      "Point data-redirect at an existing page route (data-route value or filename-derived path), or use \"/\" for the index route",
		OperationID: a.OperationID,
	}}
}
