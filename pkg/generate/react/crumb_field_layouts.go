//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what crumbFieldLayouts — data-crumb-field 페이지를 호스팅하는 레이아웃명 집합 (3단 배정 사슬 재사용, 선언 없으면 nil)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// crumbFieldLayouts returns the set of layout names hosting at least one
// data-crumb-field page (plans/stml/sitemap Phase006) — exactly the
// layouts that need the dynamic crumb-label wiring (state + pathname
// reset + Outlet context). The layout of each page follows the Phase003
// three-step assignment chain (page data-layout > sitemap nav data-layout
// > defaultLayout — the buildRoutes judgment, so wiring and route
// grouping never disagree). nil when the sitemap declares no crumb field,
// keeping every layout emission byte-identical; a layout-less page ("")
// contributes nothing — without a layout there is no breadcrumb to label.
func crumbFieldLayouts(pages []stml.PageSpec, sitemap *stml.SitemapSpec, defaultLayout string) map[string]bool {
	crumbPages := stml.SitemapCrumbFields(sitemap)
	if len(crumbPages) == 0 {
		return nil
	}
	pageLayouts := sitemapPageLayouts(sitemap)
	out := map[string]bool{}
	for _, p := range pages {
		if _, ok := crumbPages[p.Name]; !ok {
			continue
		}
		layout := p.Layout
		if layout == "" {
			layout = pageLayouts[p.Name]
		}
		if layout == "" {
			layout = defaultLayout
		}
		if layout != "" {
			out[layout] = true
		}
	}
	return out
}
