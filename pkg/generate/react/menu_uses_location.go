//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what menuUsesLocation — 조상 하이라이트 prefix 매칭이 필요한 메뉴 항목 존재 여부 (useLocation import 게이트)

package react

// menuUsesLocation reports whether any rendered menu item carries
// ancestor-highlight prefixes — the gate for the useLocation import and
// the `const { pathname } = useLocation()` line in the layout body.
func menuUsesLocation(items []sitemapMenuItem) bool {
	for _, it := range items {
		if len(it.Prefixes) > 0 || menuUsesLocation(it.Children) {
			return true
		}
	}
	return false
}
