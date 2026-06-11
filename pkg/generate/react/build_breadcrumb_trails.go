//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what buildBreadcrumbTrails — 사이트맵 nav 들을 순회하며 페이지별 정적 crumb trail 목록 구성 (generate 시점 계산)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// buildBreadcrumbTrails computes every page's static breadcrumb trail
// from the sitemap tree at generate time (plans/stml/sitemap Phase004,
// DESIGN §4.6 — trails are constants of the tree, no runtime walk).
// Trails come back in document order across all nav blocks; the per-node
// judgment (depth ≥ 2, resolvable page, ancestor href rules) lives in
// appendBreadcrumbTrails. nil sitemap → nil, keeping every breadcrumb
// emission off byte-identically.
func buildBreadcrumbTrails(sitemap *stml.SitemapSpec, routePatterns map[string]string) []breadcrumbTrail {
	if sitemap == nil {
		return nil
	}
	var trails []breadcrumbTrail
	for _, nav := range sitemap.Navs {
		appendBreadcrumbTrails(nav.Items, 1, nil, routePatterns, &trails)
	}
	return trails
}
