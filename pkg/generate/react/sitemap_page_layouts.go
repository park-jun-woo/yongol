//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what sitemapPageLayouts — 페이지명 → 소속 nav 블록 data-layout 맵 (레이아웃 배정 3단 사슬의 중간 단계)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapPageLayouts maps every page listed under a sitemap nav block that
// declares data-layout to that layout — the middle link of the Phase003
// layout assignment chain (page data-layout > sitemap block data-layout >
// defaultLayout, DESIGN §4.9). Blocks without data-layout delegate to
// defaultLayout and contribute nothing; TM-40 already forbids a page
// appearing twice, so first-wins is a formality. nil sitemap → nil map.
func sitemapPageLayouts(sitemap *stml.SitemapSpec) map[string]string {
	if sitemap == nil {
		return nil
	}
	out := map[string]string{}
	for _, nav := range sitemap.Navs {
		if nav.Layout == "" {
			continue
		}
		stack := append([]stml.SitemapNode(nil), nav.Items...)
		for len(stack) > 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if _, ok := out[n.Page]; n.Page != "" && !ok {
				out[n.Page] = nav.Layout
			}
			stack = append(stack, n.Children...)
		}
	}
	return out
}
