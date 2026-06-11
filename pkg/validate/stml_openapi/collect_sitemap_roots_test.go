//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what collectSitemapRoots — entry 블록 루트 + 메뉴 렌더 루트 동시 편입 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectSitemapRoots(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "login", FileName: "login.html"},
	}
	names := map[string]bool{"dashboard": true, "login": true}
	g := &pageGraph{Roots: map[string]bool{}, InSitemap: map[string]bool{}, MenuBlocked: map[string]string{}}
	sm := &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{
			{Items: []stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Menu: true}}},
			{Entry: true, Items: []stml.SitemapNode{{Page: "login", Label: "로그인", Menu: false}}},
		},
	}
	collectSitemapRoots(sm, names, pages, g)
	if !g.Roots["dashboard"] {
		t.Error("the menu-rendered dashboard should be a root")
	}
	if !g.Roots["login"] {
		t.Error("the entry-block login should be a root even with data-menu=\"false\"")
	}
}
