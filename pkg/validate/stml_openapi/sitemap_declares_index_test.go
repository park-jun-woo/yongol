//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestSitemapDeclaresIndex — nil/미선언 false, 중첩 data-index true 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapDeclaresIndex(t *testing.T) {
	t.Run("nil sitemap declares nothing", func(t *testing.T) {
		if sitemapDeclaresIndex(nil) {
			t.Error("nil sitemap should not declare an index")
		}
	})

	t.Run("sitemap without data-index", func(t *testing.T) {
		sm := &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "dashboard", Label: "대시보드", Menu: true},
			}}},
		}
		if sitemapDeclaresIndex(sm) {
			t.Error("expected false without any data-index entry")
		}
	})

	t.Run("nested data-index counts", func(t *testing.T) {
		sm := &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Label: "그룹", Children: []stml.SitemapNode{
					{Page: "dashboard", Label: "대시보드", Index: true, Menu: true},
				}},
			}}},
		}
		if !sitemapDeclaresIndex(sm) {
			t.Error("expected true for a nested data-index entry")
		}
	})
}
