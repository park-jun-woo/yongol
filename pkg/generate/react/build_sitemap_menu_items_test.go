//ff:func feature=gen-react type=test control=sequence
//ff:what buildSitemapMenuItems — 페이지/그룹/외부 링크 종별·아이콘 변환·prefix 수집과 중복 제거·깊이 제한·data-menu=false 서브트리 은닉 검증

package react

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildSitemapMenuItems(t *testing.T) {
	patterns := map[string]string{
		"dashboard":    "/dashboard",
		"order-list":   "/orders",
		"order-detail": "/orders/:OrderID",
		"order-edit":   "/orders/:OrderID/edit",
		"report-list":  "/reports",
	}

	t.Run("three kinds: page, group, external", func(t *testing.T) {
		nodes := []stml.SitemapNode{
			{Page: "dashboard", Label: "대시보드", Menu: true, Icon: "layout-dashboard"},
			{Label: "주문", Menu: true, Children: []stml.SitemapNode{
				{Page: "order-list", Label: "주문 목록", Menu: true},
			}},
			{Label: "도움말", Href: "https://docs.example.com", Menu: true},
		}
		items := buildSitemapMenuItems(nodes, 1, patterns)
		if len(items) != 3 {
			t.Fatalf("expected 3 items, got %d: %+v", len(items), items)
		}

		page := items[0]
		if page.Kind != "page" || page.To != "/dashboard" || page.Icon != "LayoutDashboard" {
			t.Errorf("page item = %+v", page)
		}
		if len(page.Prefixes) != 0 || len(page.Children) != 0 {
			t.Errorf("leaf page must carry no prefixes/children: %+v", page)
		}

		group := items[1]
		if group.Kind != "group" || group.To != "" || group.Href != "" || group.Icon != "" {
			t.Errorf("group item = %+v", group)
		}
		if len(group.Children) != 1 || group.Children[0].To != "/orders" {
			t.Errorf("group children = %+v", group.Children)
		}

		ext := items[2]
		if ext.Kind != "external" || ext.Href != "https://docs.example.com" || ext.To != "" {
			t.Errorf("external item = %+v", ext)
		}
	})

	t.Run("hidden descendants contribute deduped prefixes", func(t *testing.T) {
		nodes := []stml.SitemapNode{
			{Page: "order-list", Label: "주문 목록", Menu: true, Children: []stml.SitemapNode{
				// required param → menu-hidden → "/orders/" prefix
				{Page: "order-detail", Label: "상세", Menu: true},
				// same static prefix "/orders/" — must dedupe
				{Page: "order-edit", Label: "수정", Menu: true},
				// data-menu="false" → menu-hidden → static "/reports" prefix
				{Page: "report-list", Label: "리포트", Menu: false},
			}},
		}
		items := buildSitemapMenuItems(nodes, 1, patterns)
		if len(items) != 1 {
			t.Fatalf("expected 1 item, got %d: %+v", len(items), items)
		}
		if want := []string{"/orders/", "/reports"}; !reflect.DeepEqual(items[0].Prefixes, want) {
			t.Errorf("Prefixes = %v, want %v", items[0].Prefixes, want)
		}
		if len(items[0].Children) != 0 {
			t.Errorf("all children are menu-hidden, none must render: %+v", items[0].Children)
		}
	})

	t.Run("data-menu=false hides the whole subtree", func(t *testing.T) {
		nodes := []stml.SitemapNode{
			{Page: "order-list", Label: "주문 목록", Menu: false, Children: []stml.SitemapNode{
				{Page: "dashboard", Label: "대시보드", Menu: true},
			}},
		}
		if items := buildSitemapMenuItems(nodes, 1, patterns); len(items) != 0 {
			t.Errorf("hidden parent must hide renderable children too: %+v", items)
		}
	})

	t.Run("depth beyond 2 renders nothing", func(t *testing.T) {
		nodes := []stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Menu: true}}
		if items := buildSitemapMenuItems(nodes, 3, patterns); len(items) != 0 {
			t.Errorf("depth-3 nodes must not render: %+v", items)
		}
	})

	t.Run("required-param page is skipped, unresolved page falls back", func(t *testing.T) {
		nodes := []stml.SitemapNode{
			{Page: "order-detail", Label: "상세", Menu: true},
			{Page: "unknown-page", Label: "미지", Menu: true},
		}
		items := buildSitemapMenuItems(nodes, 1, patterns)
		if len(items) != 1 {
			t.Fatalf("required-param page must be skipped, got %+v", items)
		}
		if items[0].Kind != "page" || items[0].To != "/unknown-page" {
			t.Errorf("unresolved page must fall back to /<page-name>: %+v", items[0])
		}
	})
}
