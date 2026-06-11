//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what menuHasNavLink — 내부 페이지 메뉴 항목 존재 여부 (NavLink import 게이트)

package react

// menuHasNavLink reports whether any rendered menu item is an internal
// page link — the gate for the react-router-dom NavLink import (groups
// and external links alone need none). A dynamic menu group (Phase007)
// counts: its fetched items render as NavLinks.
func menuHasNavLink(items []sitemapMenuItem) bool {
	for _, it := range items {
		if it.Kind == "page" || it.Fetch != "" || menuHasNavLink(it.Children) {
			return true
		}
	}
	return false
}
