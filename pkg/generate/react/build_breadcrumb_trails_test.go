//ff:func feature=gen-react type=test control=sequence
//ff:what buildBreadcrumbTrails — 그룹 포함 trail / 조상 href 유무 / 깊이 1 생략 / 미등재·미해석 페이지 생략 / nil sitemap 검증

package react

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildBreadcrumbTrails(t *testing.T) {
	routePatterns := map[string]string{
		"dashboard":       "/dashboard",
		"building-list":   "/buildings",
		"building-detail": "/buildings/:BuildingID",
	}
	sitemap := &stml.SitemapSpec{Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
		{Page: "dashboard", Label: "대시보드", Index: true, Menu: true}, // depth 1 — no trail
		{Label: "건물 관리", Menu: true, Children: []stml.SitemapNode{ // group — label-only crumb
			{Page: "building-list", Label: "건물 목록", Menu: true, Children: []stml.SitemapNode{
				{Page: "building-detail", Label: "건물 상세", Menu: true},
			}},
			{Page: "ghost-page", Label: "유령", Menu: true}, // unresolved — skipped (TM-39's finding)
		}},
	}}}}

	trails := buildBreadcrumbTrails(sitemap, routePatterns)
	want := []breadcrumbTrail{
		{Page: "building-list", Pattern: "/buildings", Crumbs: []breadcrumbCrumb{
			{Label: "건물 관리"},
			{Label: "건물 목록"}, // self — label only despite being menu-renderable
		}},
		{Page: "building-detail", Pattern: "/buildings/:BuildingID", Crumbs: []breadcrumbCrumb{
			{Label: "건물 관리"},
			{Label: "건물 목록", Href: "/buildings"}, // menu-renderable ancestor — linked
			{Label: "건물 상세"},
		}},
	}
	if !reflect.DeepEqual(trails, want) {
		t.Errorf("trails = %+v\nwant %+v", trails, want)
	}

	if got := buildBreadcrumbTrails(nil, routePatterns); got != nil {
		t.Errorf("nil sitemap should yield nil trails, got %+v", got)
	}
}
