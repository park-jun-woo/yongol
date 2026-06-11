//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what SitemapCrumbFields — 사이트맵 전체에서 페이지명 → data-crumb-field 맵 구성 (부재 시 nil)

package stml

// SitemapCrumbFields collects every page's data-crumb-field declaration
// across the whole sitemap (plans/stml/sitemap Phase006) into a page-name
// → field map — the shared lookup of the page emitter (dynamic crumb
// effect), the layout emitter (Outlet context wiring) and the breadcrumb
// artifacts. nil sitemap or no declaration → nil, keeping every emission
// keyed off it byte-identical to the crumb-field-less output. Group <li>s
// (no page) contribute nothing — their misplaced data-crumb-field is
// TM-39's finding. First occurrence wins (a duplicate page is TM-40's).
func SitemapCrumbFields(sm *SitemapSpec) map[string]string {
	if sm == nil {
		return nil
	}
	var fields map[string]string
	for _, nav := range sm.Navs {
		fields = addSitemapCrumbFields(nav.Items, fields)
	}
	return fields
}
