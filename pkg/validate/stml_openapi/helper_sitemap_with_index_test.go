//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what sitemapWithIndexFixture — 단일 data-index 항목을 가진 SitemapSpec fixture 생성

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapWithIndexFixture builds a one-entry sitemap whose entry carries
// data-index on the named page — the shared TM-42 test fixture.
func sitemapWithIndexFixture(page string) *stml.SitemapSpec {
	return &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: page, Label: page, Index: true, Menu: true},
		}}},
	}
}
