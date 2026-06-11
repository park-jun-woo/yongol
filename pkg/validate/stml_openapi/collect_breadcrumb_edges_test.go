//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what collectBreadcrumbEdges — 자식→MenuRenderable 조상 간선 / 그룹·필수 파라미터 조상 제외 / 미실존 페이지 필터 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectBreadcrumbEdges(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "building-list", FileName: "building-list.html"},
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"},
		{Name: "unit-detail", FileName: "unit-detail.html", Route: "/units/:UnitID"},
	}
	names := map[string]bool{"building-list": true, "building-detail": true, "unit-detail": true}
	sm := &stml.SitemapSpec{Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
		{Label: "건물 관리", Menu: true, Children: []stml.SitemapNode{ // group — never a link target
			{Page: "building-list", Label: "건물 목록", Menu: true, Children: []stml.SitemapNode{
				{Page: "building-detail", Label: "건물 상세", Menu: true, Children: []stml.SitemapNode{
					// ancestor building-detail has a required param — crumb gets
					// no href, so no edge to it; building-list still gets one.
					{Page: "unit-detail", Label: "세대 상세", Menu: true},
				}},
				{Page: "ghost", Label: "유령", Menu: true}, // nonexistent — filtered (TM-39's finding)
			}},
		}},
	}}}}

	g := &pageGraph{Edges: map[string][]string{}}
	collectBreadcrumbEdges(sm, names, pages, g)

	if want := []string{"building-list"}; !reflect.DeepEqual(g.Edges["building-detail"], want) {
		t.Errorf("Edges[building-detail] = %v, want %v (menu-renderable ancestor only)", g.Edges["building-detail"], want)
	}
	if want := []string{"building-list"}; !reflect.DeepEqual(g.Edges["unit-detail"], want) {
		t.Errorf("Edges[unit-detail] = %v, want %v (required-param ancestor skipped, group skipped)", g.Edges["unit-detail"], want)
	}
	if got := g.Edges["building-list"]; got != nil {
		t.Errorf("Edges[building-list] = %v, want none (no page ancestors)", got)
	}
	if got := g.Edges["ghost"]; got != nil {
		t.Errorf("Edges[ghost] = %v, want none (nonexistent page filtered)", got)
	}
}
