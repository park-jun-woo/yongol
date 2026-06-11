//ff:func feature=gen-react type=test control=sequence
//ff:what appendBreadcrumbTrails — 필수 파라미터 조상 라벨만 / 깊이 3 초과 트리 / data-menu=false 조상 href 없음 / crumb-field 자기 crumb Dynamic 검증

package react

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAppendBreadcrumbTrails(t *testing.T) {
	routePatterns := map[string]string{
		"tool-list":   "/tools",
		"tool-detail": "/tools/:ToolID",
		"tool-runs":   "/tool-runs/:ToolID",
	}
	// tool-list hides itself from the menu (data-menu="false") → its crumb
	// carries no href; tool-detail's route needs :ToolID → label-only too.
	nodes := []stml.SitemapNode{
		{Page: "tool-list", Label: "도구 목록", Menu: false, Children: []stml.SitemapNode{
			{Page: "tool-detail", Label: "도구 상세", Menu: true, Children: []stml.SitemapNode{
				{Page: "tool-runs", Label: "실행 이력", Menu: true},
			}},
		}},
	}

	var trails []breadcrumbTrail
	appendBreadcrumbTrails(nodes, 1, nil, routePatterns, &trails)

	want := []breadcrumbTrail{
		{Page: "tool-detail", Pattern: "/tools/:ToolID", Crumbs: []breadcrumbCrumb{
			{Label: "도구 목록"}, // data-menu="false" — no href
			{Label: "도구 상세"},
		}},
		{Page: "tool-runs", Pattern: "/tool-runs/:ToolID", Crumbs: []breadcrumbCrumb{
			{Label: "도구 목록"},
			{Label: "도구 상세"}, // required-param ancestor — no href
			{Label: "실행 이력"},
		}},
	}
	if !reflect.DeepEqual(trails, want) {
		t.Errorf("trails = %+v\nwant %+v", trails, want)
	}

	// data-crumb-field marks only the trail's own crumb Dynamic (Phase006)
	// — the same page as an ancestor stays static (self label only).
	dynNodes := []stml.SitemapNode{
		{Page: "tool-list", Label: "도구 목록", Menu: true, Children: []stml.SitemapNode{
			{Page: "tool-detail", Label: "도구 상세", Menu: true, CrumbField: "tool_name", Children: []stml.SitemapNode{
				{Page: "tool-runs", Label: "실행 이력", Menu: true},
			}},
		}},
	}
	var dynTrails []breadcrumbTrail
	appendBreadcrumbTrails(dynNodes, 1, nil, routePatterns, &dynTrails)
	if len(dynTrails) != 2 {
		t.Fatalf("expected 2 trails, got %+v", dynTrails)
	}
	if self := dynTrails[0].Crumbs[1]; !self.Dynamic || self.Label != "도구 상세" {
		t.Errorf("tool-detail self crumb = %+v, want Dynamic with the static label kept", self)
	}
	if anc := dynTrails[1].Crumbs[1]; anc.Dynamic {
		t.Errorf("tool-detail as an ancestor of tool-runs must not be Dynamic: %+v", anc)
	}
}
