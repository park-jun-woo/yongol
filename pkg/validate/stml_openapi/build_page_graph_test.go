//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what buildPageGraph — 노드 전수 / 루트(인덱스·entry·메뉴 렌더) / link·redirect 간선 / 미실존 대상 제외 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildPageGraph(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "building-list", FileName: "building-list.html", Children: []stml.ChildNode{
			{Kind: "link", Link: &stml.LinkRef{TargetPage: "building-detail"}},
			{Kind: "link", Link: &stml.LinkRef{TargetPage: "ghost-page"}},
		}},
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID",
			Actions: []stml.ActionBlock{{OperationID: "DeleteBuilding", Redirect: "building-list"}}},
		{Name: "login", FileName: "login.html"},
	}
	fs := makeFS(pages, nil)
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{
			{Layout: "app", Items: []stml.SitemapNode{
				{Page: "dashboard", Label: "대시보드", Index: true, Menu: true},
				{Page: "building-list", Label: "건물", Menu: true, Children: []stml.SitemapNode{
					{Page: "building-detail", Label: "건물 상세", Menu: true},
				}},
			}},
			{Layout: "bare", Entry: true, Items: []stml.SitemapNode{
				{Page: "login", Label: "로그인", Menu: true},
			}},
		},
	}

	g := buildPageGraph(fs)
	if want := []string{"dashboard", "building-list", "building-detail", "login"}; !reflect.DeepEqual(g.Pages, want) {
		t.Errorf("Pages = %v, want %v", g.Pages, want)
	}
	for _, root := range []string{"dashboard", "building-list", "login"} {
		if !g.Roots[root] {
			t.Errorf("%q should be a root (index / menu-rendered / entry)", root)
		}
	}
	if g.Roots["building-detail"] {
		t.Error("building-detail (required param) must not be a root — listing is a node, not an edge")
	}
	if want := []string{"building-detail"}; !reflect.DeepEqual(g.Edges["building-list"], want) {
		t.Errorf("Edges[building-list] = %v, want %v (ghost target dropped)", g.Edges["building-list"], want)
	}
	if want := []string{"building-list"}; !reflect.DeepEqual(g.Edges["building-detail"], want) {
		t.Errorf("Edges[building-detail] = %v, want %v (redirect edge)", g.Edges["building-detail"], want)
	}
	if !g.InSitemap["building-detail"] || g.MenuBlocked["building-detail"] == "" {
		t.Errorf("building-detail should be listed with a menu-block reason, got InSitemap=%v reason=%q", g.InSitemap["building-detail"], g.MenuBlocked["building-detail"])
	}
}
