//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what RoutePaths — 페이지 라우트 경로 유도의 단일 소스 (data-route 우선, route.* 소비 집합 → 필수/선택 세그먼트)
package stml

import "strings"

// RoutePaths returns the route path patterns the page resolves to. It is
// the single source of route planning — react's pageToRoutes, TM-26 and
// resolveIndexRedirect all consume this table.
//
//  0. If Route is set (data-route), use it as-is (single pattern).
//  1. Base path: strip .html → "/<kebab>" ("workflows.html" → "/workflows");
//     a "-detail" suffix maps to the pluralized parent resource path
//     ("workflow-detail.html" → "/workflows").
//  2. Every route.<Name> the page consumes becomes a path segment after
//     the base: params consumed by some data-fetch are required (":Name"),
//     params consumed only by data-action blocks are optional (":Name?",
//     react-router v6.5+ optional segment) — required first, then
//     optional, each group in first-appearance order.
//  3. A page that consumes no route.* keeps the bare base path.
func RoutePaths(p PageSpec) []string {
	if p.Route != "" {
		return []string{p.Route}
	}
	base := strings.TrimSuffix(p.FileName, ".html")
	path := "/" + base
	if strings.HasSuffix(base, "-detail") {
		path = "/" + naivePluralize(strings.TrimSuffix(base, "-detail"))
	}
	params := collectRouteParams(p)
	for _, rp := range params {
		if rp.Required {
			path += "/:" + rp.Name
		}
	}
	for _, rp := range params {
		if !rp.Required {
			path += "/:" + rp.Name + "?"
		}
	}
	return []string{path}
}
