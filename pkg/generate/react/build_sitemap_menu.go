//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what buildSitemapMenu — 레이아웃 귀속 nav 블록들을 문서 순서로 연결해 메뉴 항목 모델 구성

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// buildSitemapMenu builds one layout's menu model from its sitemap nav
// blocks: the items of every block concatenate in document order into a
// single list (multiple blocks per layout chain together, DESIGN §4.9 /
// Phase003 rule 6). The judgment of what renders lives in
// buildSitemapMenuItems.
func buildSitemapMenu(navs []stml.SitemapNav, routePatterns map[string]string) []sitemapMenuItem {
	var items []sitemapMenuItem
	for _, nav := range navs {
		items = append(items, buildSitemapMenuItems(nav.Items, 1, routePatterns)...)
	}
	return items
}
