//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what menuSitemap — TM-51 테스트용 단일 nav 사이트맵 (최상위 페이지 1개가 메뉴 렌더)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// menuSitemap is a one-nav sitemap whose single top-level page renders in
// the menu (Menu true, depth 1, no required route param).
func menuSitemap() *stml.SitemapSpec {
	return &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
		{Page: "home", Label: "Home", Menu: true},
	}}}}
}
