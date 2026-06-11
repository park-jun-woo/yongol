//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what sitemapHasIcon — 사이트맵 전체에 data-icon 선언이 1건 이상 존재하는지 (lucide-react 의존성 게이트)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapHasIcon reports whether any node of the sitemap declares
// data-icon — the gate for adding the lucide-react dependency to
// package.json (the same conditional pattern as the auth-store zustand
// dependency). The whole tree is walked, not just menu-rendered nodes: a
// declared icon pulls the dependency even while its node is hidden, so
// toggling data-menu never breaks the install.
func sitemapHasIcon(sitemap *stml.SitemapSpec) bool {
	if sitemap == nil {
		return false
	}
	var stack []stml.SitemapNode
	for _, nav := range sitemap.Navs {
		stack = append(stack, nav.Items...)
	}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n.Icon != "" {
			return true
		}
		stack = append(stack, n.Children...)
	}
	return false
}
