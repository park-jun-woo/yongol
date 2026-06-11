//ff:type feature=gen-react type=model
//ff:what breadcrumbCrumb — 브레드크럼 조각 하나 (라벨 + 선택적 href; href 는 MenuRenderable 조상만)

package react

// breadcrumbCrumb is one piece of a breadcrumb trail (plans/stml/sitemap
// Phase004, DESIGN §4.6). Href is set only when the crumb's sitemap node
// is a page the shared stml_openapi.MenuRenderable judgment admits — a
// required-parameter ancestor, a group label, an external link and the
// trail's own page stay label-only (the Refine "parent without a list =
// crumb without href" semantics). Dynamic marks the trail's own crumb of
// a data-crumb-field page (Phase006): the <Breadcrumb> component renders
// the layout's runtime label state for it when set, falling back to the
// static Label — the single fallback point, so the crumb never blanks.
type breadcrumbCrumb struct {
	Label   string // sitemap <li> text (page name fallback for labelless page nodes)
	Href    string // resolved route path, "" = label-only crumb
	Dynamic bool   // data-crumb-field self crumb — runtime label slot (Phase006)
}
