//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectConsumedOps — STML data-fetch/data-action + 레이아웃 data-logout + 사이트맵 동적 그룹 data-fetch + 컴포넌트 .tsx의 api 호출에서 소비된 operationId 집합 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectConsumedOps returns the set of operationIds consumed across all STML
// pages: data-fetch and data-action blocks, layout data-logout operations
// (page-flow Phase010 — the layout emits an api.<Op>() call, so the op is a
// real consumer; without this a layout-consumed op would false-fire XMO-10
// and a stale no-front tag on it would escape XMO-12), sitemap dynamic menu
// group data-fetch declarations (plans/stml/sitemap Phase007 — the layout
// useQuery is a real consumer too), plus api.<operationId>( calls inside
// referenced components' .tsx files under <specsDir>/frontend/components/.
// Component candidates are filtered against ops (the set of real OpenAPI
// operationIds) to avoid false positives. A zero specsDir or empty ops skips
// the component scan, leaving behavior identical to the fetch/action-only
// collection.
func collectConsumedOps(pages []stml.PageSpec, layouts []stml.LayoutSpec, sitemap *stml.SitemapSpec, specsDir string, ops map[string]struct{}) map[string]struct{} {
	out := make(map[string]struct{})
	for _, page := range pages {
		for _, f := range page.Fetches {
			collectFetchOps(f, out)
		}
		for _, a := range page.Actions {
			out[a.OperationID] = struct{}{}
		}
	}
	for _, l := range layouts {
		if l.Logout != nil && l.Logout.OperationID != "" {
			out[l.Logout.OperationID] = struct{}{}
		}
	}
	if sitemap != nil {
		for _, e := range collectSitemapEntries(sitemap) {
			if e.Node.Fetch != "" {
				out[e.Node.Fetch] = struct{}{}
			}
		}
	}
	names := collectComponentNames(pages)
	collectComponentApiOps(names, specsDir, ops, out)
	return out
}
