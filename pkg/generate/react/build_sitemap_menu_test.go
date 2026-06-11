//ff:func feature=gen-react type=test control=sequence
//ff:what buildSitemapMenu — 그룹/직속 항목/2단 초과 생략/data-menu=false 서브트리 생략/필수 파라미터 생략/외부 링크/prefix/아이콘/블록 연결 검증

package react

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildSitemapMenu(t *testing.T) {
	patterns := map[string]string{
		"dashboard":          "/dashboard",
		"building-list":      "/buildings",
		"building-detail":    "/buildings/:BuildingID",
		"building-documents": "/building-documents",
		"building-rights":    "/building-rights",
		"member-list":        "/members",
		"member-detail":      "/members/:MemberID",
	}

	t.Run("representative tree", func(t *testing.T) {
		navs := []stml.SitemapNav{{Layout: "app", Items: []stml.SitemapNode{
			{Page: "dashboard", Label: "대시보드", Menu: true, Icon: "layout-dashboard"},
			{Label: "건물 관리", Menu: true, Children: []stml.SitemapNode{
				{Page: "building-list", Label: "건물 목록", Menu: true, Children: []stml.SitemapNode{
					{Page: "building-detail", Label: "건물 상세", Menu: true},
					{Page: "building-documents", Label: "문서", Menu: true},
				}},
				{Page: "building-rights", Label: "권리관계", Menu: true},
			}},
			{Page: "member-list", Label: "멤버", Menu: false, Children: []stml.SitemapNode{
				{Page: "member-detail", Label: "멤버 상세", Menu: true},
			}},
			{Label: "매뉴얼", Href: "https://docs.example.com", Menu: true},
		}}}

		items := buildSitemapMenu(navs, patterns)
		if len(items) != 3 {
			t.Fatalf("expected 3 top-level items (member-list hidden by data-menu), got %d: %+v", len(items), items)
		}

		dash := items[0]
		if dash.Kind != "page" || dash.To != "/dashboard" || dash.Icon != "LayoutDashboard" {
			t.Errorf("dashboard item = %+v", dash)
		}

		group := items[1]
		if group.Kind != "group" || group.Label != "건물 관리" || len(group.Children) != 2 {
			t.Fatalf("group item = %+v", group)
		}
		list := group.Children[0]
		if list.To != "/buildings" {
			t.Errorf("building-list To = %q", list.To)
		}
		// depth-3 children never render in the menu — they contribute
		// ancestor-highlight prefixes instead (required param and paramless alike).
		if len(list.Children) != 0 {
			t.Errorf("depth-3 nodes must not render: %+v", list.Children)
		}
		wantPrefixes := []string{"/buildings/", "/building-documents"}
		if !reflect.DeepEqual(list.Prefixes, wantPrefixes) {
			t.Errorf("Prefixes = %v, want %v", list.Prefixes, wantPrefixes)
		}
		if len(group.Children[1].Prefixes) != 0 {
			t.Errorf("leaf item must carry no prefixes: %+v", group.Children[1])
		}

		ext := items[2]
		if ext.Kind != "external" || ext.Href != "https://docs.example.com" {
			t.Errorf("external item = %+v", ext)
		}
	})

	t.Run("required-param page hidden at any depth", func(t *testing.T) {
		navs := []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "building-detail", Label: "건물 상세", Menu: true},
		}}}
		if items := buildSitemapMenu(navs, patterns); len(items) != 0 {
			t.Errorf("required-param page must not render: %+v", items)
		}
	})

	t.Run("multiple blocks concatenate in document order", func(t *testing.T) {
		navs := []stml.SitemapNav{
			{Items: []stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Menu: true}}},
			{Items: []stml.SitemapNode{{Page: "building-rights", Label: "권리관계", Menu: true}}},
		}
		items := buildSitemapMenu(navs, patterns)
		if len(items) != 2 || items[0].To != "/dashboard" || items[1].To != "/building-rights" {
			t.Errorf("concatenation broken: %+v", items)
		}
	})
}
