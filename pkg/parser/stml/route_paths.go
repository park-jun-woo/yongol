//ff:func feature=stml-parse type=util control=sequence
//ff:what RoutePaths — PageSpec이 해석되는 라우트 경로 패턴 목록 (react pageToRoutes와 동기 유지)
package stml

import "strings"

// RoutePaths returns the route path patterns the page resolves to:
//
//  0. If Route is set (data-route), use it as-is (single pattern)
//  1. Strip .html → kebab-case path (e.g. "workflows.html" → "/workflows")
//  2. "-detail" suffix → parent resource path + /:id (single pattern)
//  3. Non-detail page with a route param → base path + base/:id
//
// Mirrors pkg/generate/react/page_to_routes.go (pageToRoutes) — kept in
// sync manually until route planning is extracted to a shared layer.
func RoutePaths(p PageSpec) []string {
	if p.Route != "" {
		return []string{p.Route}
	}
	base := strings.TrimSuffix(p.FileName, ".html")
	if strings.HasSuffix(base, "-detail") {
		parent := strings.TrimSuffix(base, "-detail")
		return []string{"/" + naivePluralize(parent) + "/:id"}
	}
	routePath := "/" + base
	if pageHasRouteParam(p) {
		return []string{routePath, routePath + "/:id"}
	}
	return []string{routePath}
}
