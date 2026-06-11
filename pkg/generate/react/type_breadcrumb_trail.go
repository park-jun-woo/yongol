//ff:type feature=gen-react type=model
//ff:what breadcrumbTrail — 한 페이지의 정적 브레드크럼 (루트→자신 crumb 사슬 + 라우트 매칭 패턴)

package react

// breadcrumbTrail is the static breadcrumb of one sitemap-listed page
// (plans/stml/sitemap Phase004): the root-to-self chain of sitemap labels,
// computed at generate time — the trail is a constant of the tree, so no
// runtime tree walk is emitted. Pattern is the page's resolved route
// pattern (stml.RoutePaths first pattern, the same table App.tsx mounts);
// the generated <Breadcrumb> component matches the current pathname
// against it to select the trail. Depth-1 pages get no trail at all — a
// single-crumb breadcrumb is noise (plan rule 4).
type breadcrumbTrail struct {
	Page    string            // STML page name (BREADCRUMBS key)
	Pattern string            // resolved route pattern (BREADCRUMB_ROUTES entry)
	Crumbs  []breadcrumbCrumb // root → self, document order
}
