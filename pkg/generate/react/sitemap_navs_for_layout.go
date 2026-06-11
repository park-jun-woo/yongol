//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what sitemapNavsForLayout — 레이아웃에 귀속되는 sitemap nav 블록 선별 (data-layout 일치 또는 "" = defaultLayout 위임)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapNavsForLayout selects the sitemap nav blocks that feed one
// layout's menu, document order preserved: blocks whose data-layout names
// the layout, plus blocks without data-layout when the layout IS the
// manifest defaultLayout (the "" = defaultLayout delegation of
// SitemapNav). nil sitemap → nil (the layout keeps its data-nav path).
func sitemapNavsForLayout(sitemap *stml.SitemapSpec, layoutName, defaultLayout string) []stml.SitemapNav {
	if sitemap == nil {
		return nil
	}
	var navs []stml.SitemapNav
	for _, nav := range sitemap.Navs {
		if nav.Layout == layoutName || (nav.Layout == "" && layoutName == defaultLayout) {
			navs = append(navs, nav)
		}
	}
	return navs
}
