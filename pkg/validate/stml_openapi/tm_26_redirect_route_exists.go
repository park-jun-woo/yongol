//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-26 — data-redirect 가 정적 경로로 해석되지 않거나 페이지명 참조가 STML 페이지 집합에 없음 (ERROR, "/"는 인덱스로 허용)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm26RedirectRouteExists checks that an action's data-redirect resolves
// to an STML page. A "/"-prefixed value is a static path matched against
// every page's resolved route patterns (explicit data-route or the
// filename-derived patterns, ":param" segments matching any one segment);
// "/" is always allowed — Phase005 emits the index route. Any other value
// is a page-name reference (page-flow Phase008): it must name an existing
// STML page (filename without .html), like TM-31 for data-link. A
// redirect into the void would generate a navigate() to a route that
// renders nothing.
func tm26RedirectRouteExists(a stml.ActionBlock, file string, pages []stml.PageSpec) []diagnostic.Diagnostic {
	if a.Redirect == "" || a.Redirect == "/" {
		return nil
	}
	if !strings.HasPrefix(a.Redirect, "/") {
		for _, p := range pages {
			if p.Name == a.Redirect {
				return nil
			}
		}
		return []diagnostic.Diagnostic{{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-26] data-redirect %q on action %q does not name any STML page", a.Redirect, a.OperationID),
			Advice:      "Use the target page's STML filename without .html (a page-name reference), or a \"/\"-prefixed static path",
			OperationID: a.OperationID,
		}}
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
