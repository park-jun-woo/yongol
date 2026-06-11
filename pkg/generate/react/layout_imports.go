//ff:func feature=gen-react type=util control=sequence
//ff:what 레이아웃에 필요한 react-router-dom named import 목록을 반환한다 (sitemap 메뉴는 NavLink/useLocation, data-nav 는 Link; logout 방출 시 useNavigate)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// layoutImports returns the react-router-dom named imports needed for a
// layout. With a sitemap-derived menu (menu != nil — plans/stml/sitemap
// Phase003) the menu needs NavLink for internal page items and useLocation
// when any item carries ancestor-highlight prefixes or the dynamic crumb
// label is wired (Phase006 — the pathname-change reset reads it); the
// legacy data-nav path (menu == nil) keeps Link byte-identically.
// emitLogout reports whether the layout emits the data-logout button
// (declared AND auth present — page-flow Phase010), which needs
// useNavigate for the /login destination.
func layoutImports(layout stml.LayoutSpec, emitLogout bool, menu *sitemapMenu) []string {
	var imports []string
	if menu != nil {
		if menuHasNavLink(menu.Items) {
			imports = append(imports, "NavLink")
		}
		if layout.HasOutlet {
			imports = append(imports, "Outlet")
		}
		if menuUsesLocation(menu.Items) || (menu.DynamicCrumb && layout.HasOutlet) {
			imports = append(imports, "useLocation")
		}
		if emitLogout {
			imports = append(imports, "useNavigate")
		}
		return imports
	}
	if len(layout.NavItems) > 0 {
		imports = append(imports, "Link")
	}
	if layout.HasOutlet {
		imports = append(imports, "Outlet")
	}
	if emitLogout {
		imports = append(imports, "useNavigate")
	}
	return imports
}
