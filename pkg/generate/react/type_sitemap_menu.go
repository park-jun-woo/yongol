//ff:type feature=gen-react type=model
//ff:what sitemapMenu — 한 레이아웃의 사이트맵 파생 메뉴 (nil = sitemap 부재 → 현행 data-nav 경로)

package react

// sitemapMenu carries the sitemap-derived menu of one layout. A nil
// *sitemapMenu means frontend/sitemap.html is absent and the layout keeps
// the legacy data-nav emission path byte-identically; a non-nil value with
// zero items means the sitemap exists but assigns nothing to this layout
// (data-nav is then ignored — TM-44 rejects its coexistence upstream).
type sitemapMenu struct {
	Items []sitemapMenuItem
	// RoleField is frontend.auth.role_field — the claims[] key the
	// data-roles menu filter reads (plans/stml/sitemap Phase005). ""
	// renders every item unconditionally (no data-roles wiring).
	RoleField string
	// DynamicCrumb wires the dynamic breadcrumb label (plans/stml/sitemap
	// Phase006) — true when some page hosted by this layout declares
	// data-crumb-field: the layout keeps a crumb-label state, resets it on
	// pathname change, feeds it to <Breadcrumb label={...}> and hands the
	// setter down through <Outlet context>. false keeps the Phase004/005
	// emission byte-identical.
	DynamicCrumb bool
}
